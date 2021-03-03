package project

var (
	description string = ""
	gitSHA             = "n/a"
	name        string = "cluster-api-control-plane-app"
	source      string = "https://github.com/giantswarm/cluster-api-control-plane-app"
	version            = "0.0.0"
)

func Description() string {
	return description
}

func GitSHA() string {
	return gitSHA
}

func Name() string {
	return name
}

func Source() string {
	return source
}

func Version() string {
	return version
}
