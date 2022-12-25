package errpkg

import "picket/src/app"

type auth struct {
	CodeNotValid   *app.Error `yaml:"codeNotValid"`
	Unauthorized   *app.Error `yaml:"unauthorized"`
	AccountExisted *app.Error `yaml:"accountExisted"`
	LoginFailed    *app.Error `yaml:"loginFailed"`
}

type test struct {
	TestHasContent *app.Error `yaml:"testHasContent"`
}

type answersheet struct {
	UserNotDoing  *app.Error `yaml:"userNotDoing"`
	TimeNotValid  *app.Error `yaml:"timeNotValid"`
	UserDoingTest *app.Error `yaml:"userDoingTest"`
}

type general struct {
	Unauthorized *app.Error `yaml:"unauthorized"`
	Internal     *app.Error `yaml:"internal"`
	Forbidden    *app.Error `yaml:"forbidden"`
	BadRequest   *app.Error `yaml:"badRequest"`
	NotFound     *app.Error `yaml:"notFound"`
}

type rootErr struct {
	Auth        *auth        `yaml:"auth"`
	Test        *test        `yaml:"test"`
	Answersheet *answersheet `yaml:"answersheet"`
	General     *general     `yaml:"general"`
}
