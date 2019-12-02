#### Tontine
- Provide server to support group's tontine
- We can add investor in a tontine
- Update profit for a group. The profit will be share for all members in group

#### How to run

- Install sql-migrate

```
go get -v github.com/rubenv/sql-migrate/...
```

- Follow step-by-step to run project
```
make init
make db-migrate-up
make run
```

- Install Insomnia
- Import `tontine_insomnia.json` into Insomnia app
