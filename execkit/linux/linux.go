package linux

import "os/exec"

func Download(url, output string) *exec.Cmd {
	return exec.Command("curl", url, "-o", output)
}

func Tar(path, option, directory string) *exec.Cmd {
	return exec.Command("tar", option, path, "-C", directory)
}

func Remove(file string) *exec.Cmd {
	return exec.Command("rm", file)
}
