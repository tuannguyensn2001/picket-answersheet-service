package errpkg

import "picket/src/app"

type answersheet struct {
	NotFoundUserStart *app.Error `yaml:"notFoundUserStart"`
	UserStarted       *app.Error `yaml:"userStarted"`
	TimeNotValid      *app.Error `yaml:"timeNotValid"`
	UserDoingTest     *app.Error `yaml:"userDoingTest"`
	UserNotDoing      *app.Error `yaml:"userNotDoing"`
}

type general struct {
	NotFound     *app.Error `yaml:"notFound"`
	Unauthorized *app.Error `yaml:"unauthorized"`
	Internal     *app.Error `yaml:"internal"`
	Forbidden    *app.Error `yaml:"forbidden"`
	BadRequest   *app.Error `yaml:"badRequest"`
}

type auth struct {
	CodeNotValid   *app.Error `yaml:"codeNotValid"`
	Unauthorized   *app.Error `yaml:"unauthorized"`
	AccountExisted *app.Error `yaml:"accountExisted"`
	LoginFailed    *app.Error `yaml:"loginFailed"`
}

type test struct {
	TestHasContent *app.Error `yaml:"testHasContent"`
}

type rootErr struct {
	Answersheet *answersheet `yaml:"answersheet"`
	General     *general     `yaml:"general"`
	Auth        *auth        `yaml:"auth"`
	Test        *test        `yaml:"test"`
}
