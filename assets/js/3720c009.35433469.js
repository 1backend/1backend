"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[4787],{81478:(t,e,a)=>{a.r(e),a.d(e,{default:()=>h});a(96540);var s=a(34164),l=a(38512),n=a(90811),r=a(28400),i=a(51191),c=a(87907),o=a(24861),g=a(74848);function u(t){let{title:e}=t;return(0,g.jsxs)(g.Fragment,{children:[(0,g.jsx)(l.be,{title:e}),(0,g.jsx)(c.A,{tag:"doc_tags_list"})]})}function d(t){let{tags:e,title:a}=t;return(0,g.jsx)(l.e3,{className:(0,s.A)(n.G.page.docsTagsListPage),children:(0,g.jsx)("div",{className:"container margin-vert--lg",children:(0,g.jsx)("div",{className:"row",children:(0,g.jsxs)("main",{className:"col col--8 col--offset-2",children:[(0,g.jsx)(o.default,{as:"h1",children:a}),(0,g.jsx)(i.A,{tags:e})]})})})})}function h(t){const e=(0,r.b)();return(0,g.jsxs)(g.Fragment,{children:[(0,g.jsx)(u,{...t,title:e}),(0,g.jsx)(d,{...t,title:e})]})}},61871:(t,e,a)=>{a.d(e,{A:()=>i});a(96540);var s=a(34164),l=a(7634);const n={tag:"tag_zVej",tagRegular:"tagRegular_sFm0",tagWithCount:"tagWithCount_h2kH"};var r=a(74848);function i(t){let{permalink:e,label:a,count:i,description:c}=t;return(0,r.jsxs)(l.default,{href:e,title:c,className:(0,s.A)(n.tag,i?n.tagWithCount:n.tagRegular),children:[a,i&&(0,r.jsx)("span",{children:i})]})}},51191:(t,e,a)=>{a.d(e,{A:()=>o});a(96540);var s=a(28400),l=a(61871),n=a(24861);const r={tag:"tag_Nnez"};var i=a(74848);function c(t){let{letterEntry:e}=t;return(0,i.jsxs)("article",{children:[(0,i.jsx)(n.default,{as:"h2",id:e.letter,children:e.letter}),(0,i.jsx)("ul",{className:"padding--none",children:e.tags.map((t=>(0,i.jsx)("li",{className:r.tag,children:(0,i.jsx)(l.A,{...t})},t.permalink)))}),(0,i.jsx)("hr",{})]})}function o(t){let{tags:e}=t;const a=(0,s.Q)(e);return(0,i.jsx)("section",{className:"margin-vert--lg",children:a.map((t=>(0,i.jsx)(c,{letterEntry:t},t.letter)))})}},28400:(t,e,a)=>{a.d(e,{Q:()=>n,b:()=>l});var s=a(59979);const l=()=>(0,s.translate)({id:"theme.tags.tagsPageTitle",message:"Tags",description:"The title of the tag list page"});function n(t){const e={};return Object.values(t).forEach((t=>{const a=function(t){return t[0].toUpperCase()}(t.label);e[a]??=[],e[a].push(t)})),Object.entries(e).sort(((t,e)=>{let[a]=t,[s]=e;return a.localeCompare(s)})).map((t=>{let[e,a]=t;return{letter:e,tags:a.sort(((t,e)=>t.label.localeCompare(e.label)))}}))}}}]);