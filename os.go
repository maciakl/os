package main    

import (
"os"
"fmt"
"flag"
"bytes"
"strings"
"runtime"
"os/exec"
"path/filepath"
)

const version = "0.1.1"


func main() {

    var hideName bool
    var showIcon bool
    var showDetail bool
    var help, version bool

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
    flag.BoolVar(&version, "version", false, "Show version and quit")
    flag.BoolVar(&version, "v", false, "show version and quit")
    flag.BoolVar(&help, "help", false, "Show usage and quit")
    flag.BoolVar(&help, "h", false, "show usage and quit")
    flag.Parse()

    if version {
        Version()
        os.Exit(0)
    }

    if help {
        Usage()
        os.Exit(0)
    }

    if len(flag.Args()) > 1 {
        switch os.Args[1] {
        case "version":
            Version()
        case "help":
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
    fmt.Println("  -d, -detail     Hide the name of the operating system")
    os.Exit(0)
}


func showOperatingSystem(hideName, showIcon, showDetail bool) {
    
    name := runtime.GOOS
    icon := getIconFromName(name)
    detail := getDetail(name)

    if name == "linux" {
        icon = GetLinuxIcon(detail)
    }

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
            return ""
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
            return linuxDetail()
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
    capt, err := cmd.Output()
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }

    capt = capt[8:] // remove "Caption" from the output
    capt = bytes.TrimSpace(capt) // remove leading and trailing whitespaces

    ver, err := exec.Command("reg", "query", "HKLM\\SOFTWARE\\Microsoft\\Windows NT\\CurrentVersion", "/v", "DisplayVersion").Output()
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }

    ver = bytes.TrimSpace(ver[bytes.Index(ver, []byte("REG_SZ"))+6:])

    out := capt
    out = append(out, ' ')
    out = append(out, ver...)

    return string(out)
}

func macDetail() string {

    var product, version string

    // run sw_vers 
    cmd := exec.Command("sw_vers")
    out, err := cmd.Output()
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }

    // get the ProductName from the output
    lines := bytes.Split(out, []byte("\n"))
    for _, line := range lines {
        if bytes.HasPrefix(line, []byte("ProductName:")) {
            line = bytes.TrimPrefix(line, []byte("ProductName:"))
            line = bytes.TrimSpace(line)
             product = string(line)
        }

        if bytes.HasPrefix(line, []byte("ProductVersion:")) {
            line = bytes.TrimPrefix(line, []byte("ProductVersion:"))
            line = bytes.TrimSpace(line)
             version = string(line)
        }
    }

    output := product + " " + version
    return output
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
    // run uname -mrs

    cmd := exec.Command("uname", "-mrs")
    out, err := cmd.Output()
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }

    return string(out)
}

func GetLinuxIcon(detail string) string {
    
    // check if the detail contains "Ubuntu"
    if strings.Contains(detail, "Ubuntu") {
        return "";
    }
    
    // check if the detail contains "Debian"
    if strings.Contains(detail, "Debian") {
        return "";
    }

    // check if the detail contains Arch
    if strings.Contains(detail, "Arch") {
        return "";
    }

    // check if the detail contains Fedora
    if strings.Contains(detail, "Fedora") {
        return "";
    }

    // check if the detail contains Kali
    if strings.Contains(detail, "Kali") {
        return "";
    }

    // check if the detail contains Mint
    if strings.Contains(detail, "Mint") {
        return "󰣭";
    }

    // check if the detail contains CentOS
    if strings.Contains(detail, "CentOS") {
        return "";
    }

    // check if the detail contains Red Hat
    if strings.Contains(detail, "Red Hat") {
        return "";
    }

    // check if the detail contains SUSE
    if strings.Contains(detail, "SUSE") {
        return "";
    }

    // check if the detail contains Gentoo
    if strings.Contains(detail, "Gentoo") {
        return "";
    }

    // check if the detail contains Slackware
    if strings.Contains(detail, "Slackware") {
        return "";
    }

    // check if the detail contains Alpine
    if strings.Contains(detail, "Alpine") {
        return "";
    }

    // check if the detail contains Raspbian
    if strings.Contains(detail, "Raspbian") {
        return "";
    }

    // check if the detail contains Void
    if strings.Contains(detail, "Void") {
        return "";
    }

    // check if the detail contains Manjaro
    if strings.Contains(detail, "Manjaro") {
        return "";
    }

    // check if the detail contains openSUSE
    if strings.Contains(detail, "openSUSE") {
        return "";
    }

    // check if the detail contains elementary
    if strings.Contains(detail, "elementary") {
        return "";
    }

    // check if the detail contains Pop!_OS
    if strings.Contains(detail, "Pop!_OS") {
        return "";
    }

    // check if the detail contains Parrot
    if strings.Contains(detail, "Parrot") {
        return "";
    }

    // check if the detail contains Zorin
    if strings.Contains(detail, "Zorin") {
        return "";
    }

    // check if the detail contains NixOS
    if strings.Contains(detail, "NixOS") {
        return "";
    }

    // check if the detail contains Lubuntu
    if strings.Contains(detail, "Lubuntu") {
        return "󰕈";
    }

    // check if the detail contains Xubuntu
    if strings.Contains(detail, "Xubuntu") {
        return "󰕈";
    }


    return "";
}
