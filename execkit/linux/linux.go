package linux

import "os/exec"

func Download(url, output string) *exec.Cmd {
	return exec.Command("curl", url, "-o", output)
}

func ExtractGzip(path, directory string) *exec.Cmd {
	return exec.Command("tar", "xvzf", path, "-C", directory)
}

func Remove(file string) *exec.Cmd {
	return exec.Command("rm", file)
}

func Bash(command string) *exec.Cmd {
	return exec.Command("sh", "-c", command)
}
