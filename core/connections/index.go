package connections

import "net"

type User struct {
	UserID int64
}

type Connect struct {
	Conn net.Conn
	User User
}

var connections map[int64]Connect

func Save(user User, conn net.Conn) {
	connections[user.UserID] = Connect{Conn:conn,User:user}
}

func Drop(userID int64, conn net.Conn) {
	delete(connections, userID)
}

func Get(userId int64) net.Conn {
	return connections[userId]
}
