- Comando para rodar o Commitsar localmente:
docker run --rm --name="commitsar" -w /src -v "$(pwd)":/src aevea/commitsar

- Projeto possui Commitzen instalado para controle de commit-lint
1) Instale o commitzen globalmente com o comando:
npm install -g commitizen
2) utilize com o comando para executar o commit:
git cz

