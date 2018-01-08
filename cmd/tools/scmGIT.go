package main

import (
	"fmt"
	"time"

	"gopkg.in/src-d/go-git.v4"
	. "gopkg.in/src-d/go-git.v4/_examples"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/storage/memory"
)

//#####################
// ****func ScmGIT****
// Clones the given repository, creating the remote, the local branches
// and fetching the objects, everything in memory:
// retrieve the HEAD reference
// retrieves the commit history
// Compare time and sort all VM having 3 days old branch
//###############################

func ScmGIT(x string, y plumbing.ReferenceName) (time.Time, string) {
	// Clones the given repository, creating the remote, the local branches
	// and fetching the objects, everything in memory:
	Info("GIT Check is in progress....")
	Info("GIT CLONE")
	fmt.Println(x)
	Info("GIT BRANCH")
	fmt.Println(y)
	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		ReferenceName: y,
		URL:           x,
	})
	CheckIfError(err)
	// ... retrieving the HEAD reference
	ref, err := r.Head()
	CheckIfError(err)

	// ... retrieves the commit history
	cIter, err := r.Log(&git.LogOptions{From: ref.Hash()})
	CheckIfError(err)
	commit, err := cIter.Next()
	dateCommit := commit.Committer.When
	CheckIfError(err)

	// ... Compare time
	now := time.Now()
	diff := now.Sub(dateCommit)
	days := int(diff.Hours() / 24)
	var u string
	if days < 3 {
		u = "InUsed"
		return dateCommit, u
	}
	u = "Unused"
	return dateCommit, u

}
