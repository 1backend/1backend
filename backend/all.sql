DROP DATABASE IF EXISTS backend;
CREATE DATABASE IF NOT EXISTS backend;
USE backend;

CREATE TABLE IF NOT EXISTS projects
(
  id                  VARCHAR(36)                                                     NOT NULL,
  infra_password      VARCHAR(36)                                                     NOT NULL,
  author 	            VARCHAR(128)                                                    NOT NULL,
  description         VARCHAR(256)                                                    NOT NULL,
  name                VARCHAR(128)                                                    NOT NULL,
  public              BOOLEAN                                                         NOT NULL,
  open_source         BOOLEAN                                                         NOT NULL,
  mode                VARCHAR(32)                                                     NOT NULL,
  recipe              VARCHAR(32)                                                     NOT NULL,
  imports             TEXT                                                            NOT NULL,
  packages            TEXT                                                            NOT NULL,
  source              TEXT                                                            NOT NULL,
  port                INT                                                             NOT NULL,
  stars               INT                                                             NOT NULL,
  read_me             TEXT                                                            NOT NULL,
  created_at          DATETIME DEFAULT CURRENT_TIMESTAMP                              NOT NULL,
  updated_at          DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  NOT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

ALTER TABLE projects
      ADD CONSTRAINT uniqueProjectPerAuthor UNIQUE (author, name);

CREATE TABLE IF NOT EXISTS issues
(
  id                  VARCHAR(36)                                                     NOT NULL,
  project_id          VARCHAR(36)                                                     NOT NULL,
  title               VARCHAR(256)                                                    NOT NULL,
  comment_count       INT                                                             NOT NULL,
  user_id             VARCHAR(36)                                                     NOT NULL,
  created_at          DATETIME DEFAULT CURRENT_TIMESTAMP                              NOT NULL,
  updated_at          DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  NOT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

ALTER TABLE issues
      ADD FOREIGN KEY (project_id) REFERENCES projects (id),
      ADD FOREIGN KEY (user_id) REFERENCES users (id);

CREATE TABLE IF NOT EXISTS comments
(
  id                  VARCHAR(36)                                                     NOT NULL,
  issue_id            VARCHAR(36)                                                     NOT NULL,
  content             TEXT                                                            NOT NULL,
  user_id             VARCHAR(36)                                                     NOT NULL,
  created_at          DATETIME DEFAULT CURRENT_TIMESTAMP                              NOT NULL,
  updated_at          DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  NOT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

ALTER TABLE issues
      ADD FOREIGN KEY (issue_id) REFERENCES issues (id),
      ADD FOREIGN KEY (user_id) REFERENCES users (id);

CREATE TABLE IF NOT EXISTS dependencies
(
  id                  VARCHAR(36)                                                     NOT NULL,
  project_id          VARCHAR(36)                                                     NOT NULL,
  type                VARCHAR(36)                                                     NOT NULL,
  config              TEXT                                                            NOT NULL,
  created_at          DATETIME DEFAULT CURRENT_TIMESTAMP                              NOT NULL,
  updated_at          DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  NOT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

ALTER TABLE dependencies
      ADD CONSTRAINT uniqueDepPerProject UNIQUE (project_id, type),
      ADD FOREIGN KEY (project_id) REFERENCES projects (id);

CREATE TABLE IF NOT EXISTS endpoints
(
  id                  VARCHAR(36)                                                     NOT NULL,
  url                 VARCHAR(128)                                                    NOT NULL,
  method              VARCHAR(6)                                                      NOT NULL,
  description         VARCHAR(256)                                                    NOT NULL,
  `code` 		          TEXT							                                              NOT NULL,
  project_id          VARCHAR(36)                                                     NOT NULL,
  created_at          DATETIME DEFAULT CURRENT_TIMESTAMP                              NOT NULL,
  updated_at          DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  NOT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

ALTER TABLE endpoints
      ADD CONSTRAINT uniqueEndpointPerProject UNIQUE (project_id, method, url),
      ADD FOREIGN KEY (project_id) REFERENCES projects (id);

CREATE TABLE IF NOT EXISTS builds
(
  id                  VARCHAR(36)                                                     NOT NULL,
  output              TEXT                                                            NOT NULL,
  success             BOOLEAN                                                         NOT NULL,
  in_progress         BOOLEAN                                                         NOT NULL,
  project_id          VARCHAR(36)                                                     NOT NULL,
  created_at          DATETIME DEFAULT CURRENT_TIMESTAMP                              NOT NULL,
  updated_at          DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  NOT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

ALTER TABLE builds
      ADD FOREIGN KEY (project_id) REFERENCES projects (id);

CREATE TABLE IF NOT EXISTS users
(
  id                  VARCHAR(36)                                                     NOT NULL,
  nick                VARCHAR(32)                                                     NOT NULL,
  name                VARCHAR(32)                                                     NOT NULL,
  password            VARCHAR(128)                                                    NOT NULL,
  premium             BOOLEAN                                                         NOT NULL,
  email               VARCHAR(128)                                                    NOT NULL,
  avatar_link         VARCHAR(128)                                                    NOT NULL,
  created_at          DATETIME DEFAULT CURRENT_TIMESTAMP                              NOT NULL,
  updated_at          DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  NOT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS access_tokens
(
  id                  VARCHAR(36)                                                     NOT NULL,
  token               VARCHAR(36)                                             UNIQUE NOT NULL,
  user_id             VARCHAR(36)                                                     NOT NULL,
  created_at          DATETIME DEFAULT CURRENT_TIMESTAMP                              NOT NULL,
  updated_at          DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  NOT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

ALTER TABLE access_tokens
      ADD FOREIGN KEY (user_id) REFERENCES users (id);

CREATE TABLE IF NOT EXISTS stars
(
  id                  VARCHAR(36)                                                     NOT NULL,
  project_id          VARCHAR(36)                                                     NOT NULL,
  user_id             VARCHAR(36)                                                     NOT NULL,
  created_at          DATETIME DEFAULT CURRENT_TIMESTAMP                              NOT NULL,
  updated_at          DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  NOT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

ALTER TABLE stars
      ADD CONSTRAINT uniqueStarPerProject UNIQUE (project_id, user_id),
      ADD FOREIGN KEY (project_id) REFERENCES projects (id),
      ADD FOREIGN KEY (user_id) REFERENCES users (id);
