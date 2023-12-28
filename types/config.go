package types

type Config struct {
	ClientId        string
	ClientSecret    string
	RequestedScopes string
}

type EnvVarConfig struct {
	ClientId     string `yaml:"ClientId"`
	ClientSecret string `yaml:"ClientSecret"`
}

type CombinedTokenStructure struct {
	ModifyToken      UserModifyTokenStructure      `yaml:"ModifyToken"`
	ReadToken        UserReadTokenStructure        `yaml:"ReadToken"`
	LibraryReadToken UserLibraryReadTokenStructure `yaml:"LibraryReadToken"`
}

type UserModifyTokenStructure struct {
	UserModifyToken          string `yaml:"UserModifyToken"`
	UserModifyRefreshToken   string `yaml:"UserModifyRefreshToken"`
	UserModifyTokenExpiresIn int64  `yaml:"UserModifyTokenExpiresIn"`
}

type UserReadTokenStructure struct {
	UserReadToken          string `yaml:"UserReadToken"`
	UserReadRefreshToken   string `yaml:"UserReadRefreshToken"`
	UserReadTokenExpiresIn int64  `yaml:"UserReadTokenExpiresIn"`
}

type UserLibraryReadTokenStructure struct {
	UserLibraryReadToken          string `yaml:"UserLibraryReadToken"`
	UserLibraryReadRefreshToken   string `yaml:"UserLibraryReadRefreshToken"`
	UserLibraryReadTokenExpiresIn int64  `yaml:"UserLibraryReadTokenExpiresIn"`
}