package dbinitquery

// Init constants to create the database structure
// USERS Table
const InitUsersQuery = `CREATE TABLE users(
	"idUser" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
	"userName" TEXT NOT NULL,
	"userPassword" TEXT NOT NULL,
	"isAdmin" integer NOT NULL DEFAULT 0,
	"userSpaceQuota" integer NOT NULL DEFAULT 5,
	"created" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);

	CREATE UNIQUE INDEX "idUserIndex" ON users(idUser);
	CREATE INDEX "userNameIndex" ON users(userName);`

// SHARED_LINKS Table
const InitSharedLinksQuery = `CREATE TABLE shared_links(
	"idLink" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
	"userId" integer NOT NULL,
	"linkName" TEXT NOT NULL DEFAULT "New Link",
	"linkDirectory" TEXT NOT NULL,
	"linkShortID" TEXT NOT NULL,
	"linkCreated" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX "idLinkIndex" ON shared_links(idLink);
CREATE INDEX "linkShortIDIndex" ON shared_links(linksShortID);`

// INSERT queries
const InsertUserQuery = `INSERT INTO users(userName, userPassword, isAdmin) VALUES (?, ?, ?);`

// DELETE queries
const DeleteUserQuery = `DELETE FROM users WHERE idUser=?;`

// SELECT queries
const SelectUsersQuery = `SELECT idUser, userName, isAdmin FROM users ORDER BY created;`

// SELECT BY IDS queries
const GetUserQuery = `SELECT idUser, userName, isAdmin FROM users WHERE userName=?`
