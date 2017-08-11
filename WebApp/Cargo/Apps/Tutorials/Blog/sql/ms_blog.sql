USE [Blog]
GO
/****** Object:  Table [dbo].[blog_user]    Script Date: 04/19/2017 16:14:02 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[blog_user](
	[id] [int] NOT NULL,
	[first_name] [nvarchar](50) NOT NULL,
	[last_name] [nvarchar](50) NOT NULL,
	[email] [nvarchar](50) NOT NULL,
	[website] [nvarchar](150) NULL,
 CONSTRAINT [PK_blog_user] PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX  = OFF, STATISTICS_NORECOMPUTE  = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS  = ON, ALLOW_PAGE_LOCKS  = ON) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[blog_category]    Script Date: 04/19/2017 16:14:02 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
SET ANSI_PADDING ON
GO
CREATE TABLE [dbo].[blog_category](
	[blog_category] [int] NOT NULL,
	[name] [varchar](50) NOT NULL,
	[enabled] [bit] NOT NULL,
	[date_created] [datetime] NOT NULL,
 CONSTRAINT [PK_blog_category] PRIMARY KEY CLUSTERED 
(
	[blog_category] ASC
)WITH (PAD_INDEX  = OFF, STATISTICS_NORECOMPUTE  = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS  = ON, ALLOW_PAGE_LOCKS  = ON) ON [PRIMARY]
) ON [PRIMARY]
GO
SET ANSI_PADDING OFF
GO
/****** Object:  Table [dbo].[blog_author]    Script Date: 04/19/2017 16:14:02 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[blog_author](
	[id] [nvarchar](100) NOT NULL,
 CONSTRAINT [PK_blog_author] PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX  = OFF, STATISTICS_NORECOMPUTE  = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS  = ON, ALLOW_PAGE_LOCKS  = ON) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[blog_post]    Script Date: 04/19/2017 16:14:02 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[blog_post](
	[id] [int] NOT NULL,
	[title] [nvarchar](144) NOT NULL,
	[article] [nvarchar](max) NOT NULL,
    [thumbnail] [nvarchar](max) NOT NULL,
	[author_id] [nvarchar](100) NOT NULL,
	[featuread] [bit] NOT NULL,
	[enabled] [bit] NOT NULL,
	[comments_enabled] [bit] NOT NULL,
	[views] [int] NULL,
	[date_published] [datetime2](7) NOT NULL,
 CONSTRAINT [PK_blog_post] PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX  = OFF, STATISTICS_NORECOMPUTE  = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS  = ON, ALLOW_PAGE_LOCKS  = ON) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO
/****** Object:  Table [dbo].[blog_tag]    Script Date: 04/19/2017 16:14:02 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
SET ANSI_PADDING ON
GO
CREATE TABLE [dbo].[blog_tag](
	[id] [int] NOT NULL,
	[post_id] [int] NOT NULL,
	[tag] [varchar](50) NOT NULL,
	[tag_clean] [varchar](150) NOT NULL,
 CONSTRAINT [PK_blog_tag] PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX  = OFF, STATISTICS_NORECOMPUTE  = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS  = ON, ALLOW_PAGE_LOCKS  = ON) ON [PRIMARY]
) ON [PRIMARY]
GO
SET ANSI_PADDING OFF
GO
/****** Object:  Table [dbo].[blog_related]    Script Date: 04/19/2017 16:14:02 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[blog_related](
	[blog_post_id] [int] NOT NULL,
	[blog_related_post_id] [int] NOT NULL,
 CONSTRAINT [PK_blog_related] PRIMARY KEY CLUSTERED 
(
	[blog_post_id] ASC,
	[blog_related_post_id] ASC
)WITH (PAD_INDEX  = OFF, STATISTICS_NORECOMPUTE  = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS  = ON, ALLOW_PAGE_LOCKS  = ON) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[blog_post_to_category]    Script Date: 04/19/2017 16:14:02 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[blog_post_to_category](
	[category_id] [int] NOT NULL,
	[post_id] [int] NOT NULL,
 CONSTRAINT [PK_blog_post_to_category] PRIMARY KEY CLUSTERED 
(
	[category_id] ASC,
	[post_id] ASC
)WITH (PAD_INDEX  = OFF, STATISTICS_NORECOMPUTE  = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS  = ON, ALLOW_PAGE_LOCKS  = ON) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[blog_comment]    Script Date: 04/19/2017 16:14:02 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[blog_comment](
	[id] [int] NOT NULL,
	[post_id] [int] NOT NULL,
	[is_reply_to_id] [int] NOT NULL,
	[comment] [text] NOT NULL,
	[user_id] [int] NOT NULL,
	[mark_read] [bit] NOT NULL,
	[enabled] [bit] NOT NULL,
	[date] [datetime] NOT NULL,
 CONSTRAINT [PK_blog_comment] PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX  = OFF, STATISTICS_NORECOMPUTE  = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS  = ON, ALLOW_PAGE_LOCKS  = ON) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO
/****** Object:  Default [DF_blog_category_enabled]    Script Date: 04/19/2017 16:14:02 ******/
ALTER TABLE [dbo].[blog_category] ADD  CONSTRAINT [DF_blog_category_enabled]  DEFAULT ((0)) FOR [enabled]
GO
/****** Object:  Default [DF_blog_category_date_created]    Script Date: 04/19/2017 16:14:02 ******/
ALTER TABLE [dbo].[blog_category] ADD  CONSTRAINT [DF_blog_category_date_created]  DEFAULT (getutcdate()) FOR [date_created]
GO
/****** Object:  Default [DF_blog_comment_mark_read]    Script Date: 04/19/2017 16:14:02 ******/
ALTER TABLE [dbo].[blog_comment] ADD  CONSTRAINT [DF_blog_comment_mark_read]  DEFAULT ((0)) FOR [mark_read]
GO
/****** Object:  Default [DF_blog_comment_enabled]    Script Date: 04/19/2017 16:14:02 ******/
ALTER TABLE [dbo].[blog_comment] ADD  CONSTRAINT [DF_blog_comment_enabled]  DEFAULT ((0)) FOR [enabled]
GO
/****** Object:  Default [DF_blog_comment_date]    Script Date: 04/19/2017 16:14:02 ******/
ALTER TABLE [dbo].[blog_comment] ADD  CONSTRAINT [DF_blog_comment_date]  DEFAULT (getutcdate()) FOR [date]
GO
/****** Object:  Default [DF_blog_post_featuread]    Script Date: 04/19/2017 16:14:02 ******/
ALTER TABLE [dbo].[blog_post] ADD  CONSTRAINT [DF_blog_post_featuread]  DEFAULT ((0)) FOR [featuread]
GO
/****** Object:  Default [DF_blog_post_enabled]    Script Date: 04/19/2017 16:14:02 ******/
ALTER TABLE [dbo].[blog_post] ADD  CONSTRAINT [DF_blog_post_enabled]  DEFAULT ((0)) FOR [enabled]
GO
/****** Object:  Default [DF_blog_post_comments_enabled]    Script Date: 04/19/2017 16:14:02 ******/
ALTER TABLE [dbo].[blog_post] ADD  CONSTRAINT [DF_blog_post_comments_enabled]  DEFAULT ((1)) FOR [comments_enabled]
GO
/****** Object:  Default [DF_blog_post_views]    Script Date: 04/19/2017 16:14:02 ******/
ALTER TABLE [dbo].[blog_post] ADD  CONSTRAINT [DF_blog_post_views]  DEFAULT ((0)) FOR [views]
GO
/****** Object:  ForeignKey [FK_blog_comment_blog_post]    Script Date: 04/19/2017 16:14:02 ******/
ALTER TABLE [dbo].[blog_comment]  WITH CHECK ADD  CONSTRAINT [FK_blog_comment_blog_post] FOREIGN KEY([post_id])
REFERENCES [dbo].[blog_post] ([id])
ON DELETE CASCADE
GO
ALTER TABLE [dbo].[blog_comment] CHECK CONSTRAINT [FK_blog_comment_blog_post]
GO
/****** Object:  ForeignKey [FK_blog_comment_blog_user]    Script Date: 04/19/2017 16:14:02 ******/
ALTER TABLE [dbo].[blog_comment]  WITH CHECK ADD  CONSTRAINT [FK_blog_comment_blog_user] FOREIGN KEY([user_id])
REFERENCES [dbo].[blog_user] ([id])
GO
ALTER TABLE [dbo].[blog_comment] CHECK CONSTRAINT [FK_blog_comment_blog_user]
GO
/****** Object:  ForeignKey [FK_blog_post_blog_author]    Script Date: 04/19/2017 16:14:02 ******/
ALTER TABLE [dbo].[blog_post]  WITH CHECK ADD  CONSTRAINT [FK_blog_post_blog_author] FOREIGN KEY([author_id])
REFERENCES [dbo].[blog_author] ([id])
ON UPDATE CASCADE
ON DELETE CASCADE
GO
ALTER TABLE [dbo].[blog_post] CHECK CONSTRAINT [FK_blog_post_blog_author]
GO
/****** Object:  ForeignKey [FK_blog_post_to_category_blog_category]    Script Date: 04/19/2017 16:14:02 ******/
ALTER TABLE [dbo].[blog_post_to_category]  WITH CHECK ADD  CONSTRAINT [FK_blog_post_to_category_blog_category] FOREIGN KEY([category_id])
REFERENCES [dbo].[blog_category] ([blog_category])
ON DELETE CASCADE
GO
ALTER TABLE [dbo].[blog_post_to_category] CHECK CONSTRAINT [FK_blog_post_to_category_blog_category]
GO
/****** Object:  ForeignKey [FK_blog_post_to_category_blog_post]    Script Date: 04/19/2017 16:14:02 ******/
ALTER TABLE [dbo].[blog_post_to_category]  WITH CHECK ADD  CONSTRAINT [FK_blog_post_to_category_blog_post] FOREIGN KEY([post_id])
REFERENCES [dbo].[blog_post] ([id])
ON DELETE CASCADE
GO
ALTER TABLE [dbo].[blog_post_to_category] CHECK CONSTRAINT [FK_blog_post_to_category_blog_post]
GO
/****** Object:  ForeignKey [FK_blog_related_blog_post]    Script Date: 04/19/2017 16:14:02 ******/
ALTER TABLE [dbo].[blog_related]  WITH CHECK ADD  CONSTRAINT [FK_blog_related_blog_post] FOREIGN KEY([blog_post_id])
REFERENCES [dbo].[blog_post] ([id])
GO
ALTER TABLE [dbo].[blog_related] CHECK CONSTRAINT [FK_blog_related_blog_post]
GO
/****** Object:  ForeignKey [FK_blog_related_blog_post1]    Script Date: 04/19/2017 16:14:02 ******/
ALTER TABLE [dbo].[blog_related]  WITH CHECK ADD  CONSTRAINT [FK_blog_related_blog_post1] FOREIGN KEY([blog_related_post_id])
REFERENCES [dbo].[blog_post] ([id])
GO
ALTER TABLE [dbo].[blog_related] CHECK CONSTRAINT [FK_blog_related_blog_post1]
GO
/****** Object:  ForeignKey [FK_blog_tag_blog_post]    Script Date: 04/19/2017 16:14:02 ******/
ALTER TABLE [dbo].[blog_tag]  WITH CHECK ADD  CONSTRAINT [FK_blog_tag_blog_post] FOREIGN KEY([post_id])
REFERENCES [dbo].[blog_post] ([id])
ON DELETE CASCADE
GO
ALTER TABLE [dbo].[blog_tag] CHECK CONSTRAINT [FK_blog_tag_blog_post]
GO
