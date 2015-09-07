package git

func Pull(projPath string, origin string, branch string) {
	cmdArgs := []string{"--git-dir=" + projPath + "/.git", "pull", origin, branch}
	RunCmd("git", cmdArgs)
	RunCmd("git", cmdArgs)
}
