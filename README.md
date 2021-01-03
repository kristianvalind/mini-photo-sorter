# Mini Photo Sorter (mps)
A small tool for sorting photos using EXIF data.

Run `make` to build.

Modify `droplet.applescript` to set default options, and then run `make droplet` to generate a self contained AppleScript droplet for use with ImageCapture, containing the mps binary as well.

Usage info:
```
Mini Photo Sorter (mps) by Kristian Valind
Usage: ./mps [flags] files
  -d    dry run, simulate operations but don't move any files
  -h    show usage
  -o dir
        dir in which to place output files (default ".")
  -p pattern
        the pattern for the renamed files, using golang time package formatting. The string {filename} is replaced with the original file name. (default "2006-01-02/2006-01-02_{filename}")
  -r    recurse into subdirectories of provided directories
  -s    stop when encountering an error, instead of skipping to next file
```