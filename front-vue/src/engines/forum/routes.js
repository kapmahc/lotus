import ArticlesIndex from './articles/Index'
import ArticlesNew from './articles/New'
import ArticlesEdit from './articles/Edit'

import TagsIndex from './tags/Index'
import TagsNew from './tags/New'
import TagsEdit from './tags/Edit'

import CommentsIndex from './comments/Index'
import CommentsNew from './comments/New'
import CommentsEdit from './comments/Edit'

export default [
  {name: 'forum.comments.index', path: '/shop/comments', component: CommentsIndex},
  {name: 'forum.comments.new', path: '/shop/comments/new', component: CommentsNew},
  {name: 'forum.comments.edit', path: '/shop/comments/:id/edit', component: CommentsEdit},

  {name: 'forum.tags.index', path: '/shop/tags', component: TagsIndex},
  {name: 'forum.tags.new', path: '/shop/tags/new', component: TagsNew},
  {name: 'forum.tags.edit', path: '/shop/tags/:id/edit', component: TagsEdit},

  {name: 'forum.articles.index', path: '/shop/articles', component: ArticlesIndex},
  {name: 'forum.articles.new', path: '/shop/articles/new', component: ArticlesNew},
  {name: 'forum.articles.edit', path: '/shop/articles/:id/edit', component: ArticlesEdit}
]
