"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[3515],{7994:(e,i,s)=>{s.r(i),s.d(i,{assets:()=>o,contentTitle:()=>l,default:()=>h,frontMatter:()=>t,metadata:()=>n,toc:()=>a});const n=JSON.parse('{"id":"built-in-services/config-svc","title":"Config Svc","description":"The Config Svc stores public, non-sensitive and end-user-facing data.","source":"@site/docs/built-in-services/config-svc.md","sourceDirName":"built-in-services","slug":"/built-in-services/config-svc","permalink":"/docs/built-in-services/config-svc","draft":false,"unlisted":false,"editUrl":"https://github.com/1backend/1backend/tree/main/docs-source/docs/built-in-services/config-svc.md","tags":[{"inline":true,"label":"config-svc","permalink":"/docs/tags/config-svc"},{"inline":true,"label":"configuration","permalink":"/docs/tags/configuration"},{"inline":true,"label":"services","permalink":"/docs/tags/services"}],"version":"current","sidebarPosition":80,"frontMatter":{"sidebar_position":80,"tags":["config-svc","configuration","services"]},"sidebar":"tutorialSidebar","previous":{"title":"Policy Svc","permalink":"/docs/built-in-services/policy-svc"},"next":{"title":"Model Svc","permalink":"/docs/built-in-services/model-svc"}}');var c=s(74848),r=s(28453);const t={sidebar_position:80,tags:["config-svc","configuration","services"]},l="Config Svc",o={},a=[{value:"Access Rules",id:"access-rules",level:2},{value:"Read",id:"read",level:3},{value:"Write",id:"write",level:3},{value:"Related",id:"related",level:2}];function d(e){const i={a:"a",blockquote:"blockquote",code:"code",h1:"h1",h2:"h2",h3:"h3",header:"header",li:"li",p:"p",pre:"pre",ul:"ul",...(0,r.R)(),...e.components};return(0,c.jsxs)(c.Fragment,{children:[(0,c.jsx)(i.header,{children:(0,c.jsx)(i.h1,{id:"config-svc",children:"Config Svc"})}),"\n",(0,c.jsx)(i.p,{children:"The Config Svc stores public, non-sensitive and end-user-facing data."}),"\n",(0,c.jsxs)(i.blockquote,{children:["\n",(0,c.jsxs)(i.p,{children:["This page provides a high-level overview of ",(0,c.jsx)(i.code,{children:"Config Svc"}),". For detailed information, refer to the ",(0,c.jsx)(i.a,{href:"/docs/1backend/get-config",children:"Secret Svc API documentation"}),"."]}),"\n"]}),"\n",(0,c.jsxs)(i.p,{children:["The Config Svc is less critical than it might seem\u2014most configuration happens internally through the ",(0,c.jsx)(i.a,{href:"/docs/built-in-services/secret-svc",children:"Secret Svc"}),"."]}),"\n",(0,c.jsx)(i.p,{children:"At the moment, it functions more like a minimal feature-flag service."}),"\n",(0,c.jsx)(i.h2,{id:"access-rules",children:"Access Rules"}),"\n",(0,c.jsx)(i.h3,{id:"read",children:"Read"}),"\n",(0,c.jsx)(i.p,{children:"All configs are publicly readable even without a user account."}),"\n",(0,c.jsx)(i.h3,{id:"write",children:"Write"}),"\n",(0,c.jsxs)(i.p,{children:["Any logged in user can write to the config, but only to the key that matches their slug, ie. if a user's slug is ",(0,c.jsx)(i.code,{children:"jane-doe"}),":"]}),"\n",(0,c.jsx)(i.pre,{children:(0,c.jsx)(i.code,{children:'{\n  "jane-doe": {"janesConfig": 5},\n  "someOtherKey": "hi"\n}\n'})}),"\n",(0,c.jsxs)(i.p,{children:["Only the key ",(0,c.jsx)(i.code,{children:"jane-doe"})," will be written to the ",(0,c.jsx)(i.code,{children:"Config"}),", all other keys (such as ",(0,c.jsx)(i.code,{children:"someOtherKey"}),") will be ignored."]}),"\n",(0,c.jsx)(i.h2,{id:"related",children:"Related"}),"\n",(0,c.jsxs)(i.ul,{children:["\n",(0,c.jsxs)(i.li,{children:[(0,c.jsx)(i.a,{href:"/docs/built-in-services/secret-svc",children:"Secret Svc"})," to store sensitive data like internal configuration and secrets"]}),"\n"]})]})}function h(e={}){const{wrapper:i}={...(0,r.R)(),...e.components};return i?(0,c.jsx)(i,{...e,children:(0,c.jsx)(d,{...e})}):d(e)}}}]);