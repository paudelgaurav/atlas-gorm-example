-- Modify "user_roles" table
ALTER TABLE `user_roles` ADD INDEX `fk_users_role` (`user_id`), DROP FOREIGN KEY `fk_user_roles_user`, ADD CONSTRAINT `fk_users_role` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION;
