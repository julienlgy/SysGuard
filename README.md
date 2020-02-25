# SysGuard

### Step01 - Proxy

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

