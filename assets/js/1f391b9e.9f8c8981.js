"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[6061],{10294:(e,t,n)=>{n.d(t,{A:()=>c});n(96540);var a=n(34164),s=n(32960),i=n(90811),l=n(64362),r=n(74848);function d(e){let{className:t}=e;return(0,r.jsx)(l.A,{type:"caution",title:(0,r.jsx)(s.Rc,{}),className:(0,a.A)(t,i.G.common.unlistedBanner),children:(0,r.jsx)(s.Uh,{})})}function c(e){return(0,r.jsxs)(r.Fragment,{children:[(0,r.jsx)(s.AE,{}),(0,r.jsx)(d,{...e})]})}},807:(e,t,n)=>{n.r(t),n.d(t,{default:()=>o});n(96540);var a=n(34164),s=n(32960),i=n(90811),l=n(64362),r=n(74848);function d(e){let{className:t}=e;return(0,r.jsx)(l.A,{type:"caution",title:(0,r.jsx)(s.Yh,{}),className:(0,a.A)(t,i.G.common.draftBanner),children:(0,r.jsx)(s.TT,{})})}var c=n(10294);function o(e){let{metadata:t}=e;const{unlisted:n,frontMatter:a}=t;return(0,r.jsxs)(r.Fragment,{children:[(n||a.unlisted)&&(0,r.jsx)(c.A,{}),a.draft&&(0,r.jsx)(d,{})]})}},40315:(e,t,n)=>{n.d(t,{A:()=>v});n(96540);var a=n(34164),s=n(59979),i=n(90811),l=n(7634);const r={iconEdit:"iconEdit_Z9Sw"};var d=n(74848);function c(e){let{className:t,...n}=e;return(0,d.jsx)("svg",{fill:"currentColor",height:"20",width:"20",viewBox:"0 0 40 40",className:(0,a.A)(r.iconEdit,t),"aria-hidden":"true",...n,children:(0,d.jsx)("g",{children:(0,d.jsx)("path",{d:"m34.5 11.7l-3 3.1-6.3-6.3 3.1-3q0.5-0.5 1.2-0.5t1.1 0.5l3.9 3.9q0.5 0.4 0.5 1.1t-0.5 1.2z m-29.5 17.1l18.4-18.5 6.3 6.3-18.4 18.4h-6.3v-6.2z"})})})}function o(e){let{editUrl:t}=e;return(0,d.jsxs)(l.default,{to:t,className:i.G.common.editThisPage,children:[(0,d.jsx)(c,{}),(0,d.jsx)(s.default,{id:"theme.common.editThisPage",description:"The link label to edit the current page",children:"Edit this page"})]})}var u=n(34542);function m(e){let{lastUpdatedAt:t}=e;const n=new Date(t),a=(0,u.i)({day:"numeric",month:"short",year:"numeric",timeZone:"UTC"}).format(n);return(0,d.jsx)(s.default,{id:"theme.lastUpdated.atDate",description:"The words used to describe on which date a page has been last updated",values:{date:(0,d.jsx)("b",{children:(0,d.jsx)("time",{dateTime:n.toISOString(),itemProp:"dateModified",children:a})})},children:" on {date}"})}function h(e){let{lastUpdatedBy:t}=e;return(0,d.jsx)(s.default,{id:"theme.lastUpdated.byUser",description:"The words used to describe by who the page has been last updated",values:{user:(0,d.jsx)("b",{children:t})},children:" by {user}"})}function f(e){let{lastUpdatedAt:t,lastUpdatedBy:n}=e;return(0,d.jsxs)("span",{className:i.G.common.lastUpdated,children:[(0,d.jsx)(s.default,{id:"theme.lastUpdated.lastUpdatedAtBy",description:"The sentence used to display when a page has been last updated, and by who",values:{atDate:t?(0,d.jsx)(m,{lastUpdatedAt:t}):"",byUser:n?(0,d.jsx)(h,{lastUpdatedBy:n}):""},children:"Last updated{atDate}{byUser}"}),!1]})}const x={lastUpdated:"lastUpdated_JAkA"};function v(e){let{className:t,editUrl:n,lastUpdatedAt:s,lastUpdatedBy:i}=e;return(0,d.jsxs)("div",{className:(0,a.A)("row",t),children:[(0,d.jsx)("div",{className:"col",children:n&&(0,d.jsx)(o,{editUrl:n})}),(0,d.jsx)("div",{className:(0,a.A)("col",x.lastUpdated),children:(s||i)&&(0,d.jsx)(f,{lastUpdatedAt:s,lastUpdatedBy:i})})]})}},22812:(e,t,n)=>{n.d(t,{A:()=>N});var a=n(96540),s=n(28453),i=n(50344),l=n(52398),r=n(74848);function d(e){return(0,r.jsx)("code",{...e})}var c=n(7634);var o=n(94688);function u(e){const t=a.Children.toArray(e.children),n=t.find((e=>a.isValidElement(e)&&"summary"===e.type)),s=(0,r.jsx)(r.Fragment,{children:t.filter((e=>e!==n))});return(0,r.jsx)(o.default,{...e,summary:n,children:s})}var m=n(24861);function h(e){return(0,r.jsx)(m.default,{...e})}var f=n(34164);const x={containsTaskList:"containsTaskList_mC6p"};function v(e){if(void 0!==e)return(0,f.A)(e,e?.includes("contains-task-list")&&x.containsTaskList)}var p=n(39269);const g={img:"img_ev3q"};var j=n(64362);const A={Head:i.A,details:u,Details:u,code:function(e){return function(e){return void 0!==e.children&&a.Children.toArray(e.children).every((e=>"string"==typeof e&&!e.includes("\n")))}(e)?(0,r.jsx)(d,{...e}):(0,r.jsx)(l.default,{...e})},a:function(e){return(0,r.jsx)(c.default,{...e})},pre:function(e){return(0,r.jsx)(r.Fragment,{children:e.children})},ul:function(e){return(0,r.jsx)("ul",{...e,className:v(e.className)})},li:function(e){return(0,p.A)().collectAnchor(e.id),(0,r.jsx)("li",{...e})},img:function(e){return(0,r.jsx)("img",{decoding:"async",loading:"lazy",...e,className:(t=e.className,(0,f.A)(t,g.img))});var t},h1:e=>(0,r.jsx)(h,{as:"h1",...e}),h2:e=>(0,r.jsx)(h,{as:"h2",...e}),h3:e=>(0,r.jsx)(h,{as:"h3",...e}),h4:e=>(0,r.jsx)(h,{as:"h4",...e}),h5:e=>(0,r.jsx)(h,{as:"h5",...e}),h6:e=>(0,r.jsx)(h,{as:"h6",...e}),admonition:j.A,mermaid:()=>null};function N(e){let{children:t}=e;return(0,r.jsx)(s.x,{components:A,children:t})}},50591:(e,t,n)=>{n.r(t),n.d(t,{default:()=>h});n(96540);var a=n(34164),s=n(38512),i=n(90811),l=n(29100),r=n(22812),d=n(39149),c=n(807),o=n(40315);const u={mdxPageWrapper:"mdxPageWrapper_j9I6"};var m=n(74848);function h(e){const{content:t}=e,{metadata:n,assets:h}=t,{title:f,editUrl:x,description:v,frontMatter:p,lastUpdatedBy:g,lastUpdatedAt:j}=n,{keywords:A,wrapperClassName:N,hide_table_of_contents:y}=p,L=h.image??p.image,b=!!(x||j||g);return(0,m.jsx)(s.e3,{className:(0,a.A)(N??i.G.wrapper.mdxPages,i.G.page.mdxPage),children:(0,m.jsxs)(l.A,{children:[(0,m.jsx)(s.be,{title:f,description:v,keywords:A,image:L}),(0,m.jsx)("main",{className:"container container--fluid margin-vert--lg",children:(0,m.jsxs)("div",{className:(0,a.A)("row",u.mdxPageWrapper),children:[(0,m.jsxs)("div",{className:(0,a.A)("col",!y&&"col--8"),children:[(0,m.jsx)(c.default,{metadata:n}),(0,m.jsx)("article",{children:(0,m.jsx)(r.A,{children:(0,m.jsx)(t,{})})}),b&&(0,m.jsx)(o.A,{className:(0,a.A)("margin-top--sm",i.G.pages.pageFooterEditMetaRow),editUrl:x,lastUpdatedAt:j,lastUpdatedBy:g})]}),!y&&t.toc.length>0&&(0,m.jsx)("div",{className:"col col--2",children:(0,m.jsx)(d.A,{toc:t.toc,minHeadingLevel:p.toc_min_heading_level,maxHeadingLevel:p.toc_max_heading_level})})]})})]})})}},39149:(e,t,n)=>{n.d(t,{A:()=>c});n(96540);var a=n(34164),s=n(41751);const i={tableOfContents:"tableOfContents_bqdL",docItemContainer:"docItemContainer_F8PC"};var l=n(74848);const r="table-of-contents__link toc-highlight",d="table-of-contents__link--active";function c(e){let{className:t,...n}=e;return(0,l.jsx)("div",{className:(0,a.A)(i.tableOfContents,"thin-scrollbar",t),children:(0,l.jsx)(s.A,{...n,linkClassName:r,linkActiveClassName:d})})}},41751:(e,t,n)=>{n.d(t,{A:()=>u});var a=n(96540),s=n(79762),i=n(46751),l=n(48698),r=n(7634),d=n(74848);function c(e){let{toc:t,className:n,linkClassName:a,isChild:s}=e;return t.length?(0,d.jsx)("ul",{className:s?void 0:n,children:t.map((e=>(0,d.jsxs)("li",{children:[(0,d.jsx)(r.default,{to:`#${e.id}`,className:a??void 0,dangerouslySetInnerHTML:{__html:e.value}}),(0,d.jsx)(c,{isChild:!0,toc:e.children,className:n,linkClassName:a})]},e.id)))}):null}const o=a.memo(c);function u(e){let{toc:t,className:n="table-of-contents table-of-contents__left-border",linkClassName:r="table-of-contents__link",linkActiveClassName:c,minHeadingLevel:u,maxHeadingLevel:m,...h}=e;const f=(0,s.p)(),x=u??f.tableOfContents.minHeadingLevel,v=m??f.tableOfContents.maxHeadingLevel,p=(0,i.h)({toc:t,minHeadingLevel:x,maxHeadingLevel:v}),g=(0,a.useMemo)((()=>{if(r&&c)return{linkClassName:r,linkActiveClassName:c,minHeadingLevel:x,maxHeadingLevel:v}}),[r,c,x,v]);return(0,l.i)(g),(0,d.jsx)(o,{toc:p,className:n,linkClassName:r,...h})}},48698:(e,t,n)=>{n.d(t,{i:()=>d});var a=n(96540),s=n(79762);function i(e){const t=e.getBoundingClientRect();return t.top===t.bottom?i(e.parentNode):t}function l(e,t){let{anchorTopOffset:n}=t;const a=e.find((e=>i(e).top>=n));if(a){return function(e){return e.top>0&&e.bottom<window.innerHeight/2}(i(a))?a:e[e.indexOf(a)-1]??null}return e[e.length-1]??null}function r(){const e=(0,a.useRef)(0),{navbar:{hideOnScroll:t}}=(0,s.p)();return(0,a.useEffect)((()=>{e.current=t?0:document.querySelector(".navbar").clientHeight}),[t]),e}function d(e){const t=(0,a.useRef)(void 0),n=r();(0,a.useEffect)((()=>{if(!e)return()=>{};const{linkClassName:a,linkActiveClassName:s,minHeadingLevel:i,maxHeadingLevel:r}=e;function d(){const e=function(e){return Array.from(document.getElementsByClassName(e))}(a),d=function(e){let{minHeadingLevel:t,maxHeadingLevel:n}=e;const a=[];for(let s=t;s<=n;s+=1)a.push(`h${s}.anchor`);return Array.from(document.querySelectorAll(a.join()))}({minHeadingLevel:i,maxHeadingLevel:r}),c=l(d,{anchorTopOffset:n.current}),o=e.find((e=>c&&c.id===function(e){return decodeURIComponent(e.href.substring(e.href.indexOf("#")+1))}(e)));e.forEach((e=>{!function(e,n){n?(t.current&&t.current!==e&&t.current.classList.remove(s),e.classList.add(s),t.current=e):e.classList.remove(s)}(e,e===o)}))}return document.addEventListener("scroll",d),document.addEventListener("resize",d),d(),()=>{document.removeEventListener("scroll",d),document.removeEventListener("resize",d)}}),[e,n])}},46751:(e,t,n)=>{n.d(t,{h:()=>r,v:()=>i});var a=n(96540);function s(e){const t=e.map((e=>({...e,parentIndex:-1,children:[]}))),n=Array(7).fill(-1);t.forEach(((e,t)=>{const a=n.slice(2,e.level);e.parentIndex=Math.max(...a),n[e.level]=t}));const a=[];return t.forEach((e=>{const{parentIndex:n,...s}=e;n>=0?t[n].children.push(s):a.push(s)})),a}function i(e){return(0,a.useMemo)((()=>s(e)),[e])}function l(e){let{toc:t,minHeadingLevel:n,maxHeadingLevel:a}=e;return t.flatMap((e=>{const t=l({toc:e.children,minHeadingLevel:n,maxHeadingLevel:a});return function(e){return e.level>=n&&e.level<=a}(e)?[{...e,children:t}]:t}))}function r(e){let{toc:t,minHeadingLevel:n,maxHeadingLevel:i}=e;return(0,a.useMemo)((()=>l({toc:s(t),minHeadingLevel:n,maxHeadingLevel:i})),[t,n,i])}}}]);