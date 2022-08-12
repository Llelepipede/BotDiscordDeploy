package gitmanage

import (
	"fmt"
	"time"

	"github.com/apex/log"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

const gitDataDirectory = "./ApiData/"
const defaultRemoteName = "origin"

// Commit creates a commit in the current repository
func Push(comment string) bool {
	repo, err := git.PlainOpen(gitDataDirectory)

	if err != nil {
		// Repository does not exist yet, create it
		log.Info("The Git repository does not exist yet and will be created.")

		repo, err = git.PlainInit(gitDataDirectory, false)
	}

	if err != nil {
		log.Warn("The data folder could not be converted into a Git repository. Therefore, the versioning does not work as expected.")
		return false
	}

	w, _ := repo.Worktree()

	log.Info("Committing new changes...")
	w.Add("api.json")
	w.Commit(comment, &git.CommitOptions{
		Author: &object.Signature{
			Name:  "Llelepipede",
			Email: "pchesneau3103@gmail.com",
			When:  time.Now(),
		},
	})
	fmt.Println(w.Status())

	_, err = repo.Remote(defaultRemoteName)
	if err != nil {
		if err != nil {
			log.WithError(err).Warn("Error creating remote")
		}
	}

	auth := &http.BasicAuth{
		Username: "Llelepipede",
		Password: "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",
	}
	log.Info("Pushing changes to remote")
	err = repo.Push(&git.PushOptions{
		RemoteName: defaultRemoteName,
		Auth:       auth,
	})

	if err != nil {
		log.WithError(err).Warn("Error during push")
		return false
	}
	return true
}

func Commit(comment string) bool {

	repo, err := git.PlainOpen(gitDataDirectory)

	if err != nil {
		// Repository does not exist yet, create it
		log.Info("The Git repository does not exist yet and will be created.")

		repo, err = git.PlainInit(gitDataDirectory, false)
	}

	if err != nil {
		log.Warn("The data folder could not be converted into a Git repository. Therefore, the versioning does not work as expected.")
		return false
	}

	w, _ := repo.Worktree()

	log.Info("Committing new changes...")
	w.Add("api.json")
	w.Commit(comment, &git.CommitOptions{
		Author: &object.Signature{
			Name:  "Llelepipede",
			Email: "pchesneau3103@gmail.com",
			When:  time.Now(),
		},
	})
	fmt.Println(w.Status())

	return true
}

func Pull() (*git.Repository, bool) {

	repo, err := git.PlainOpen(gitDataDirectory)
	if err != nil {
		// Repository does not exist yet, create it
		log.Info("The Git repository does not exist yet and will be created.")
		// Filesystem abstraction based on memory

		repo, err := git.PlainClone("./ApiData/", false, &git.CloneOptions{
			URL:        "https://github.com/Mentor-Paris/jsonApiGuildbot",
			RemoteName: "origin",
		})
		if err != nil {
			return repo, false
		}
	} else {

		w, _ := repo.Worktree()
		err = w.Pull(&git.PullOptions{
			RemoteName: defaultRemoteName,
		})
		if err != nil {

			log.Info("echec de pull")
			return repo, true
		}
		log.Info("reussite de pull")
	}
	return repo, true
}
