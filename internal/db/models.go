// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"database/sql"
)

type Node struct {
	ID            int64
	CreatedAt     Time
	LastSeenAt    Time
	LastInfoReqAt Time
	LastInfoResAt Time
	PublicKey     *PublicKey
	Fqdn          sql.NullString
	Motd          sql.NullString
	Version       sql.NullInt64
}

type NodeAddress struct {
	ID         int64
	CreatedAt  Time
	LastSeenAt Time
	LastPingAt Time
	LastPongAt Time
	NodeID     int64
	Net        string
	Ip         string
	Port       int64
	Ptr        sql.NullString
}
