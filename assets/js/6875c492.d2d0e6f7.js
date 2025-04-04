"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[4813],{70525:(e,t,a)=>{a.d(t,{A:()=>i});a(96540);var s=a(59979),n=a(9002),r=a(74848);function i(e){const{metadata:t}=e,{previousPage:a,nextPage:i}=t;return(0,r.jsxs)("nav",{className:"pagination-nav","aria-label":(0,s.translate)({id:"theme.blog.paginator.navAriaLabel",message:"Blog list page navigation",description:"The ARIA label for the blog pagination"}),children:[a&&(0,r.jsx)(n.A,{permalink:a,title:(0,r.jsx)(s.default,{id:"theme.blog.paginator.newerEntries",description:"The label used to navigate to the newer blog posts page (previous page)",children:"Newer entries"})}),i&&(0,r.jsx)(n.A,{permalink:i,title:(0,r.jsx)(s.default,{id:"theme.blog.paginator.olderEntries",description:"The label used to navigate to the older blog posts page (next page)",children:"Older entries"}),isNext:!0})]})}},55240:(e,t,a)=>{a.d(t,{A:()=>i});a(96540);var s=a(40509),n=a(63972),r=a(74848);function i(e){let{items:t,component:a=n.A}=e;return(0,r.jsx)(r.Fragment,{children:t.map((e=>{let{content:t}=e;return(0,r.jsx)(s.in,{content:t,children:(0,r.jsx)(a,{children:(0,r.jsx)(t,{})})},t.metadata.permalink)}))})}},87209:(e,t,a)=>{a.r(t),a.d(t,{default:()=>f});a(96540);var s=a(34164),n=a(59979),r=a(38512),i=a(90811),l=a(76641),o=a(7634),g=a(86838),c=a(70525),u=a(87907),d=a(55240),h=a(10294),p=a(24861),m=a(74848);function x(e){let{tag:t}=e;const a=(0,l.ZD)(t);return(0,m.jsxs)(m.Fragment,{children:[(0,m.jsx)(r.be,{title:a,description:t.description}),(0,m.jsx)(u.A,{tag:"blog_tags_posts"})]})}function b(e){let{tag:t,items:a,sidebar:s,listMetadata:r}=e;const i=(0,l.ZD)(t);return(0,m.jsxs)(g.A,{sidebar:s,children:[t.unlisted&&(0,m.jsx)(h.A,{}),(0,m.jsxs)("header",{className:"margin-bottom--xl",children:[(0,m.jsx)(p.default,{as:"h1",children:i}),t.description&&(0,m.jsx)("p",{children:t.description}),(0,m.jsx)(o.default,{href:t.allTagsPath,children:(0,m.jsx)(n.default,{id:"theme.tags.tagsPageLink",description:"The label of the link targeting the tag list page",children:"View All Tags"})})]}),(0,m.jsx)(d.A,{items:a}),(0,m.jsx)(c.A,{metadata:r})]})}function f(e){return(0,m.jsxs)(r.e3,{className:(0,s.A)(i.G.wrapper.blogPages,i.G.page.blogTagPostListPage),children:[(0,m.jsx)(x,{...e}),(0,m.jsx)(b,{...e})]})}},10294:(e,t,a)=>{a.d(t,{A:()=>g});a(96540);var s=a(34164),n=a(32960),r=a(90811),i=a(64362),l=a(74848);function o(e){let{className:t}=e;return(0,l.jsx)(i.A,{type:"caution",title:(0,l.jsx)(n.Rc,{}),className:(0,s.A)(t,r.G.common.unlistedBanner),children:(0,l.jsx)(n.Uh,{})})}function g(e){return(0,l.jsxs)(l.Fragment,{children:[(0,l.jsx)(n.AE,{}),(0,l.jsx)(o,{...e})]})}},76641:(e,t,a)=>{a.d(t,{Y4:()=>u,ZD:()=>l,np:()=>c,uz:()=>g,wI:()=>o});a(96540);var s=a(59979),n=a(84365),r=a(74848);function i(){const{selectMessage:e}=(0,n.W)();return t=>e(t,(0,s.translate)({id:"theme.blog.post.plurals",description:'Pluralized label for "{count} posts". Use as much plural forms (separated by "|") as your language support (see https://www.unicode.org/cldr/cldr-aux/charts/34/supplemental/language_plural_rules.html)',message:"One post|{count} posts"},{count:t}))}function l(e){const t=i();return(0,s.translate)({id:"theme.blog.tagTitle",description:"The title of the page for a blog tag",message:'{nPosts} tagged with "{tagName}"'},{nPosts:t(e.count),tagName:e.label})}function o(e){const t=i();return(0,s.translate)({id:"theme.blog.author.pageTitle",description:"The title of the page for a blog author",message:"{authorName} - {nPosts}"},{nPosts:t(e.count),authorName:e.name||e.key})}const g=()=>(0,s.translate)({id:"theme.blog.authorsList.pageTitle",message:"Authors",description:"The title of the authors page"});function c(){return(0,r.jsx)(s.default,{id:"theme.blog.authorsList.viewAll",description:"The label of the link targeting the blog authors page",children:"View all authors"})}function u(){return(0,r.jsx)(s.default,{id:"theme.blog.author.noPosts",description:"The text for authors with 0 blog post",children:"This author has not written any posts yet."})}}}]);