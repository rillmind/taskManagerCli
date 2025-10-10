# Task Manager CLI

Um gerenciador de tarefas rápidas via linha de comando escrito em Golang.

## Requisitos

- Go 1.24.8 ou superior
- git

## Instalação

Linux:

```sh
git clone https://github.com/rillmind/taskManagerCli.git
cd taskManagerCli
go build -o task-cli ./*/*.go
sudo cp task-cli /usr/bin/
```

Windows: 

```sh
git clone https://github.com/rillmind/taskManagerCli.git
cd taskManagerCli
go build -o task-cli ./*/*.go
```

obs: Se você estiver no windows, precisa adicionar o binário executável (.exe) às variávies de
ambiente manualmente. No linux basta digitar sua senha de super usuário. Caso queira utilizar sem
adicionar às variáveis de ambiente, no Linux: basta excluir a última linha e rodar './task-cli', no
Windows: basta executar normalmente com './task-cli' sem adicionar ou retirar linhas do comando.

## Informações sobre o código

### Atributos

- id: int
- description: string
- status: enum ou map com 'todo', 'in progress' e 'done'
- createdAt
- updatedAt

### Funções

- **createTask :** Criar uma task e gerar um json para armazenar a task.
- **updateTask :** Buscar task pelo id e alterar o 'title' ou 'description' pelo '.json'.
- **deleteTask :** Buscar task pelo id e excluir do '.json'.
- **listTasks :** Buscar no '.json' todas as tasks e mostrar na tela. (Pensando em deixar isso
aparecendo o tempo todo)
- **listTasksDone :** Busca no '.json' apenas as taskas marcadas como 'done'
- **listTasksNotDone :** Busca no '.json' apenas as taskas marcadas como 'not done'
- **listTasksInProgress :** Busca no '.json' apenas as taskas marcadas como 'in progress'
