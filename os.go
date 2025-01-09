package main    

import (
"os"
"fmt"
"flag"
"bytes"
"runtime"
"os/exec"
"path/filepath"
)

const version = "0.1.0"


func main() {

    var hideName bool
    var showIcon bool
    var showDetail bool

    flag.Usage = func() {
        Usage()
        os.Exit(1)
    }

    flag.BoolVar(&hideName, "name", false, "Hide the name of the operating system")
    flag.BoolVar(&hideName, "n", false, "Hide the name of the operating system")
    flag.BoolVar(&showIcon, "icon", false, "Show the icon of the operating system")
    flag.BoolVar(&showIcon, "i", false, "Show the icon of the operating system")
    flag.BoolVar(&showDetail, "detail", false, "Show the detail of the operating system")
    flag.BoolVar(&showDetail, "d", false, "Show the detail of the operating system")
    flag.Parse()

    if len(flag.Args()) > 1 {
        switch os.Args[1] {
        case "-v", "-version":
            Version()
        case "-h", "-help":
            Usage()
        default:
            Usage()
        } 
    } else {
        showOperatingSystem(hideName, showIcon, showDetail)
    }

}

func Version() {
    fmt.Println(filepath.Base(os.Args[0]), "version", version)
    os.Exit(0)
}

func Usage() {
    fmt.Println("Usage:", filepath.Base(os.Args[0]), "[options]")
    fmt.Println("Options:")
    fmt.Println("  -v, -version    Print version information and exit")
    fmt.Println("  -h, -help       Print this message and exit")
    fmt.Println("  -i, -icon       Show the icon of the operating system")
    fmt.Println("  -n, -name       Hide the name of the operating system")
    os.Exit(0)
}


func showOperatingSystem(hideName, showIcon, showDetail bool) {
    
    name := runtime.GOOS
    icon := getIconFromName(name)
    detail := getDetail(name)

    var output string

    if showIcon {
        output += icon + " "
    }

    if !hideName {
        output += name
    }

    if showDetail {
        if !hideName { output += " " }
        output += detail
    }
    
    fmt.Println(output)
}


func getIconFromName(name string) string {

    switch name {
        case "windows":
            return ""
        case "darwin":
            return ""
        case "linux":
            return ""
        case "freebsd":
            return ""
        case "netbsd":
            return ""
        case "openbsd":
            return ""
        case "solaris":
            return ""
        default:
            return "󰆧"
    }
}


func getDetail(name string) string {

    switch name {
        case "windows":
            return windowsDetail()
        case "darwin":
            return macDetail()
        case "linux":
            return linuxDetail()
        case "freebsd":
            return unixDetail()
        case "netbsd":
            return unixDetail()
        case "openbsd":
            return unixDetail()
        case "solaris":
            return unixDetail()
        default:
            return "Unknown"
    }

}

func windowsDetail() string {

    // run wmic os get caption
    cmd := exec.Command("wmic", "os", "get", "caption") 
    out, err := cmd.Output()
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }

    out = out[8:] // remove "Caption" from the output
    out = bytes.TrimSpace(out) // remove leading and trailing whitespaces

    return string(out)
}

func macDetail() string {
    // run sw_vers 
    // or /usr/lib/PlistBuddy -c "Print :ProductVersion" /System/Library/CoreServices/SystemVersion.plist
    return "Apple macOS"
}

func linuxDetail() string {

    output := "Unknown"

    // get the PRETTY_NAME from /etc/os-release
    cmd := exec.Command("cat", "/etc/os-release")
    out, err := cmd.Output()
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }

    lines := bytes.Split(out, []byte("\n"))
    for _, line := range lines {
        if bytes.HasPrefix(line, []byte("PRETTY_NAME=")) {
            line = bytes.TrimPrefix(line, []byte("PRETTY_NAME="))
            line = bytes.Trim(line, "\"")
            output = string(line)
        }       
    }

    return output
}

func unixDetail() string {
    // run uname -a

    cmd := exec.Command("uname", "-a")
    out, err := cmd.Output()
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }

    return string(out)
}
