SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='TRADITIONAL';


-- -----------------------------------------------------
-- Table `blog`.`blog_category`
-- -----------------------------------------------------
CREATE  TABLE IF NOT EXISTS `Blog`.`blog_category` (
  `id` INT NOT NULL AUTO_INCREMENT ,
  `name` VARCHAR(45) NOT NULL ,
  `name_clean` VARCHAR(45) NULL ,
  `enabled` TINYINT(1)  NOT NULL DEFAULT 1 ,
  `date_created` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ,
  PRIMARY KEY (`id`) ,
  UNIQUE INDEX `index2` (`name_clean` ASC) )
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `blog`.`blog_author`
-- -----------------------------------------------------
CREATE  TABLE IF NOT EXISTS `Blog`.`blog_author` (
  `id` INT NOT NULL AUTO_INCREMENT ,
  `display_name` VARCHAR(45) NOT NULL ,
  `first_name` VARCHAR(45) NULL ,
  `last_name` VARCHAR(45) NULL ,
  PRIMARY KEY (`id`) ,
  UNIQUE INDEX `display_name_UNIQUE` (`display_name` ASC) )
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `blog`.`blog_post`
-- -----------------------------------------------------
CREATE  TABLE IF NOT EXISTS `Blog`.`blog_post` (
  `id` INT NOT NULL AUTO_INCREMENT ,
  `title` VARCHAR(144) NOT NULL ,
  `article` TEXT NULL ,
  `title_clean` VARCHAR(144) NULL ,
  `file` VARCHAR(45) NULL ,
  `author_id` INT NOT NULL ,
  `date_published` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ,
  `banner_image` VARCHAR(144) NULL ,
  `featured` TINYINT(1)  NOT NULL DEFAULT 0 ,
  `enabled` TINYINT(1)  NOT NULL DEFAULT 0 ,
  `comments_enabled` TINYINT(1)  NOT NULL DEFAULT 1 ,
  `views` INT NOT NULL DEFAULT 0 ,
  PRIMARY KEY (`id`) ,
  INDEX `fk_post_1` (`author_id` ASC) ,
  UNIQUE INDEX `index3` (`title_clean` ASC) ,
  CONSTRAINT `fk_post_1`
    FOREIGN KEY (`author_id` )
    REFERENCES `blog`.`blog_author` (`id` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `blog`.`blog_post_to_category`
-- -----------------------------------------------------
CREATE  TABLE IF NOT EXISTS `Blog`.`blog_post_to_category` (
  `category_id` INT NOT NULL ,
  `post_id` INT NOT NULL ,
  PRIMARY KEY (`category_id`, `post_id`) ,
  INDEX `fk_posts_to_categories_1` (`category_id` ASC) ,
  INDEX `fk_posts_to_categories_2` (`post_id` ASC) ,
  CONSTRAINT `fk_posts_to_categories_1`
    FOREIGN KEY (`category_id` )
    REFERENCES `blog`.`blog_category` (`id` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_posts_to_categories_2`
    FOREIGN KEY (`post_id` )
    REFERENCES `blog`.`blog_post` (`id` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `blog`.`blog_user`
-- -----------------------------------------------------
CREATE  TABLE IF NOT EXISTS `Blog`.`blog_user` (
  `id` INT NOT NULL AUTO_INCREMENT ,
  `name` VARCHAR(45) NULL ,
  `email` VARCHAR(45) NULL ,
  `website` VARCHAR(45) NULL ,
  PRIMARY KEY (`id`) ,
  UNIQUE INDEX `email_UNIQUE` (`email` ASC) )
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `blog`.`blog_comment`
-- -----------------------------------------------------
CREATE  TABLE IF NOT EXISTS `Blog`.`blog_comment` (
  `id` INT NOT NULL AUTO_INCREMENT ,
  `post_id` INT NOT NULL ,
  `is_reply_to_id` INT NOT NULL DEFAULT 0 ,
  `comment` TEXT NOT NULL ,
  `user_id` INT NOT NULL ,
  `mark_read` TINYINT(1)  NULL DEFAULT 0 ,
  `enabled` TINYINT(1)  NOT NULL DEFAULT 1 ,
  `date` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ,
  PRIMARY KEY (`id`, `user_id`) ,
  INDEX `fk_comment_1` (`post_id` ASC) ,
  INDEX `fk_comment_2` (`user_id` ASC) ,
  CONSTRAINT `fk_comment_1`
    FOREIGN KEY (`post_id` )
    REFERENCES `blog`.`blog_post` (`id` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_comment_2`
    FOREIGN KEY (`user_id` )
    REFERENCES `blog`.`blog_user` (`id` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `blog`.`blog_tag`
-- -----------------------------------------------------
CREATE  TABLE IF NOT EXISTS `Blog`.`blog_tag` (
  `id` INT NOT NULL AUTO_INCREMENT ,
  `post_id` INT NOT NULL ,
  `tag` VARCHAR(45) NOT NULL ,
  `tag_clean` VARCHAR(45) NOT NULL ,
  PRIMARY KEY (`id`) ,
  INDEX `fk_tags_1` (`post_id` ASC) ,
  INDEX `index3` (`tag_clean` ASC) ,
  CONSTRAINT `fk_tags_1`
    FOREIGN KEY (`post_id` )
    REFERENCES `blog`.`blog_post` (`id` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `blog`.`images`
-- -----------------------------------------------------
CREATE  TABLE IF NOT EXISTS `Blog`.`images` (
  `id` INT NOT NULL AUTO_INCREMENT ,
  PRIMARY KEY (`id`) )
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `blog`.`blog_related`
-- -----------------------------------------------------
CREATE  TABLE IF NOT EXISTS `Blog`.`blog_related` (
  `blog_post_id` INT NOT NULL ,
  `blog_related_post_id` INT NOT NULL ,
  PRIMARY KEY (`blog_post_id`, `blog_related_post_id`) ,
  INDEX `fk_blog_related_1` (`blog_post_id` ASC) ,
  CONSTRAINT `fk_blog_related_1`
    FOREIGN KEY (`blog_post_id` )
    REFERENCES `blog`.`blog_post` (`id` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;



SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
