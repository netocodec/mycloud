package dbinitquery

const UsersQuery = `CREATE TABLE users(
	"idUser" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
	"userName" TEXT NOT NULL,
	"userPassword" TEXT NOT NULL,
	"isAdmin" integer NOT NULL DEFAULT 0,
	"userSpaceQuota" integer NOT NULL DEFAULT 5,
	"created" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`

const InsertUserQuery = `INSERT INTO users(userName, userPassword, isAdmin) VALUES (?, ?, ?);`

const SelectUsersQuery = `SELECT idUser, userName, isAdmin FROM users ORDER BY created`
