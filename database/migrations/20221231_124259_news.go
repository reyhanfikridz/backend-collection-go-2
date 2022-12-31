package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type News_20221231_124259 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &News_20221231_124259{}
	m.Created = "20221231_124259"

	migration.Register("News_20221231_124259", m)
}

// Run the migrations
func (m *News_20221231_124259) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE news(`id` int(11) NOT NULL AUTO_INCREMENT,`title` varchar(128) NOT NULL,`content` longtext  NOT NULL,`is_published` tinyint(1) NOT NULL,`published_at` datetime NULL,`created_at` datetime NOT NULL,`updated_at` datetime NOT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *News_20221231_124259) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `news`")
}
