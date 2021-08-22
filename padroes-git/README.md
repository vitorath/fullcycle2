# Git Glow

# GPG

Listar as chaves
gpg --list-secret-key --keyid-form LONG

Criar chave
gpg --full-generate-key

Obter public key
gpg --armor --export <key-id>

Configurar no projeto do github/outros

git config --global user.signingkey <key-id>

Commit assinado automatico
git config --global commit.gpgsign true
git config --global tag.gpgSign true

Commit assinado manual
git commit -S -m '...' 

Adicionar no bash
export GPG_TTY=$(tty)

git log --show-signature -1

Executar agent
gpcconf --launch gpg-agent
vim ~/.gnupg/gpp.conf

gpg --edit-key <key-id>
$: adduid #preenchar que pedir
$: uid <identificacao> # set o id para modificar atributos
$: trust #modificar confianca
$: save # salvar modificacoes

# Protegendo branchs
Bloquear commits main/master e develop 

# Pull Request
Pull Request - Solicitar a revisao do codigo e fazer o merge posteriormente
- Settings > Branchs > Require pull request before merging

# Criar temlates Pull Request
https://embeddedartistry.com/blog/2017/08/11/a-github-pull-request-template-for-the-ccc-process/

Criar arquivo
.github/PULL_REQUEST_TEMPLATE.md


# Code Owners
- Settings > Branchs > Require pull request before merging > Require Review from Code Owners

.github/CODEOWNERS
  *.js @usuario2
  .github/ @usuario1
  *.go @usuairo1 @grupo-x @grupo-y ..

# Semantical versioning
https://semver.org/lang/pt-BR/ **Estudar e resumir
MAJOR.MINOR.PATH => 2.1.4
MAJOR - API PUBLICA
MINOR - Adicionados funcionalidades, mas compativel com a API
PATCH - Bugs, ajustes

MAJOR = 0 - API Instavel. Pode mudar a qualquer momento.

2.1.4-alpha < 2.1.4 "alpha" (pode ser qualquer nome) mais instavel que a release

# Conventional Commits
https://www.conventionalcommits.org/en/v1.0.0/ **Estudar e resumir

# Commitlint
https://commitlint.js.org/#/

# Commitsar
Informar o local que esta querendo que verifique 
$ docker .... commitsar .

# Commitizen
https://github.com/commitizen/cz-cli
$ git cz
 
Refatorar este projeto utilizando conventional commits