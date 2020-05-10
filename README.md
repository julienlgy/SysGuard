# SysGuard

### Pre-requisites

Fetch the certificate and update the path in the file `./src/sysguard/certs.json`.

### Step01 - Proxy

`sysguard --origin=http://127.0.0.1:80 --listen=http://0.0.0.0:8080`  
`sysguard --origin=http://whiteagent.fr:80 --listen=http://0.0.0.0:80`
`export SYSPROXY=http://0.0.0.0:80 & export SYSGATEWAY=http://127.0.0.1:8080 & sysguard`

- Permet d'étabilr un pont entre le serveur web et le client. Il est possible d'établir le pont à partir d'un serveur web en local, ou alors d'une autre adresse IP (WAF stocké sur une autre machine)
- Doit pouvoir logger toute les requêtes
- Doit supporter le HTTP et le HTTPS en gérant les certificats 
- Possibilité de gérer de multiples certificats en fonction de l'host appelé. (VH)

```
GET /  HTTP 1/1
Host : google.com

Réponse : a

GET / HTTP 1/1
Host : mail.google.com

Réponse b

```
- Configuration simple à partir d'un fichier **default.conf** à la racine du projet

#### Langage utilisé 
![](https://upload.wikimedia.org/wikipedia/commons/thumb/2/23/Go_Logo_Aqua.svg/1200px-Go_Logo_Aqua.svg.png)

export GOPATH=$(pwd)
