package simplezip

import (
    "archive/zip"
    "io"
//    "strings"
//    "fmt"
//    "log"
    "os"
    "path/filepath"
)

type Simplezip struct {

  target string

}

func NewSimpleZip() *Simplezip{

  sz := new(Simplezip)

  return sz

}

func (sz *Simplezip) ZipDir(source, target string) error {
    // 1. Create a ZIP file and zip.Writer
    f, err := os.Create(target)
    if err != nil {
        return err
    }
    defer f.Close()

    writer := zip.NewWriter(f)
    defer writer.Close()

    // 2. Go through all the files of the source
    return filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        // 3. Create a local file header
        header, err := zip.FileInfoHeader(info)
        if err != nil {
            return err
        }

        // set compression
        header.Method = zip.Deflate

        // 4. Set relative path of a file as the header name
        header.Name, err = filepath.Rel(filepath.Dir(source), path)
        if err != nil {
            return err
        }
        if info.IsDir() {
            header.Name += "/"
        }

        // 5. Create writer for the file header and save content of the file
        headerWriter, err := writer.CreateHeader(header)
        if err != nil {
            return err
        }

        if info.IsDir() {
            return nil
        }

        f, err := os.Open(path)
        if err != nil {
            return err
        }
        defer f.Close()

        _, err = io.Copy(headerWriter, f)
        return err
    })
}


// Unzip will decompress a zip archived file,
// copying all files and folders
// within the zip file (parameter 1)
// to an output directory (parameter 2).

func (sz *Simplezip) UnZipDir(src string, destination string) ([]string, error) {
	
	// a variable that will store any
	//file names available in a array of strings
	var filenames []string

	// OpenReader will open the Zip file
	// specified by name and return a ReadCloser
	// Readcloser closes the Zip file,
	// rendering it unusable for I/O
	// It returns two values:
	// 1. a pointer value to ReadCloser
	// 2. an error message (if any)
	r, err := zip.OpenReader(src)

	// if there is any error then
	// (err!=nill) becomes true
	if err != nil {
		// and this block will break the loop
		// and return filenames gathered so far
		// with an err message, and move
		// back to the main function
		
		return filenames, err
	}
	
	defer r.Close()
	// defer makes sure the file is closed
	// at the end of the program no matter what.

	for _, f := range r.File {
	
	// this loop will run until there are
	// files in the source directory & will
	// keep storing the filenames and then
	// extracts into destination folder until an err arises
	
		// Store "path/filename" for returning and using later on
		fpath := filepath.Join(destination, f.Name)

//fmt.Println("fpath: ", destination)
		
/*
		// Checking for any invalid file paths
		if !strings.HasPrefix(fpath, filepath.Clean(destination)+string(os.PathSeparator)){
//fmt.Println("os.PathSeparator: ", string(os.PathSeparator))

			return filenames, fmt.Errorf("%s is an illegal filepath", fpath)
		}
*/
		
		// the filename that is accessed is now appended
		// into the filenames string array with its path
		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {
		
			// Creating a new Folder
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Creating the files in the target directory
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return filenames, err
		}

		// The created file will be stored in
		// outFile with permissions to write &/or truncate
		outFile, err := os.OpenFile(fpath,
									os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
									f.Mode())

		// again if there is any error this block
		// will be executed and process
		// will return to main function
		if err != nil {
			// with filenames gathered so far
			// and err message
			return filenames, err
		}

		rc, err := f.Open()
		
		// again if there is any error this block
		// will be executed and process
		// will return to main function		
		if err != nil {
			// with filenames gathered so far
			// and err message back to main function
			return filenames, err
		}

		_, err = io.Copy(outFile, rc)

		// Close the file without defer so that
		// it closes the outfile before the loop
		// moves to the next iteration. this kinda
		// saves an iteration of memory & time in
		// the worst case scenario.
		outFile.Close()
		rc.Close()

		// again if there is any error this block
		// will be executed and process
		// will return to main function
		if err != nil {
			// with filenames gathered so far
			// and err message back to main function
			return filenames, err
		}
	}
	
	// Finally after every file has been appended
	// into the filenames string[] and all the
	// files have been extracted into the
	// target directory, we return filenames
	// and nil as error value as the process executed
	// successfully without any errors*
	// *only if it reaches until here.
	return filenames, nil
}

