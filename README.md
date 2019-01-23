## Example queries
### test api recommend to use Insomnia
#### login  url=/login/graphql
```
mutation func($arg: LoginByGraphql!){
  login(input:$arg)
}
# query variables
{
	"arg": {
		"ID": 1,
		"password": "123456"
	}
}
```
#### get user info url=/query
```
query funcgetuser($arg: UserArg!){
  getUser(input:$arg){
    id
    name
    mail
    books{
      id
      title
      author
    }
  }
}
# query variables
{
	"arg": {
		"ID": 1
	}
}
```