"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[1601],{58581:(e,i,s)=>{s.r(i),s.d(i,{assets:()=>d,contentTitle:()=>r,default:()=>p,frontMatter:()=>a,metadata:()=>l,toc:()=>t});const l=JSON.parse('{"id":"built-in-services/file-svc","title":"File Svc","description":"The File Service handles file-related operations, including downloading files from the internet (to cache them for faster access), accepting file uploads, and serving both downloaded and uploaded files.","source":"@site/docs/built-in-services/file-svc.md","sourceDirName":"built-in-services","slug":"/built-in-services/file-svc","permalink":"/docs/built-in-services/file-svc","draft":false,"unlisted":false,"editUrl":"https://github.com/1backend/1backend/tree/main/docs-source/docs/built-in-services/file-svc.md","tags":[{"inline":true,"label":"file-svc","permalink":"/docs/tags/file-svc"},{"inline":true,"label":"file","permalink":"/docs/tags/file"},{"inline":true,"label":"containers","permalink":"/docs/tags/containers"},{"inline":true,"label":"services","permalink":"/docs/tags/services"},{"inline":true,"label":"uploads","permalink":"/docs/tags/uploads"},{"inline":true,"label":"downloads","permalink":"/docs/tags/downloads"},{"inline":true,"label":"file serve","permalink":"/docs/tags/file-serve"},{"inline":true,"label":"file serving","permalink":"/docs/tags/file-serving"},{"inline":true,"label":"file proxy","permalink":"/docs/tags/file-proxy"},{"inline":true,"label":"non-distributed","permalink":"/docs/tags/non-distributed"}],"version":"current","sidebarPosition":60,"frontMatter":{"sidebar_position":60,"tags":["file-svc","file","containers","services","uploads","downloads","file serve","file serving","file proxy","non-distributed"]},"sidebar":"tutorialSidebar","previous":{"title":"Deploy Svc","permalink":"/docs/built-in-services/deploy-svc"},"next":{"title":"Policy Svc","permalink":"/docs/built-in-services/policy-svc"}}');var n=s(74848),o=s(28453);const a={sidebar_position:60,tags:["file-svc","file","containers","services","uploads","downloads","file serve","file serving","file proxy","non-distributed"]},r="File Svc",d={},t=[{value:"Responsibilities",id:"responsibilities",level:2},{value:"Use cases",id:"use-cases",level:2},{value:"Internal file transfer",id:"internal-file-transfer",level:3},{value:"Application file uploads",id:"application-file-uploads",level:3},{value:"Component dependencies",id:"component-dependencies",level:3}];function c(e){const i={a:"a",blockquote:"blockquote",code:"code",h1:"h1",h2:"h2",h3:"h3",header:"header",li:"li",p:"p",ul:"ul",...(0,o.R)(),...e.components};return(0,n.jsxs)(n.Fragment,{children:[(0,n.jsx)(i.header,{children:(0,n.jsx)(i.h1,{id:"file-svc",children:"File Svc"})}),"\n",(0,n.jsx)(i.p,{children:"The File Service handles file-related operations, including downloading files from the internet (to cache them for faster access), accepting file uploads, and serving both downloaded and uploaded files."}),"\n",(0,n.jsx)(i.p,{children:"The File Svc is distributed. Downloads and uploads may reside on any node in the system, any File Svc node will be able to proxy them to you."}),"\n",(0,n.jsxs)(i.blockquote,{children:["\n",(0,n.jsxs)(i.p,{children:["This page provides a high-level overview of ",(0,n.jsx)(i.code,{children:"File Svc"}),". For detailed information, refer to the ",(0,n.jsx)(i.a,{href:"/docs/1backend/download-file",children:"File Svc API documentation"}),"."]}),"\n"]}),"\n",(0,n.jsx)(i.h2,{id:"responsibilities",children:"Responsibilities"}),"\n",(0,n.jsxs)(i.ul,{children:["\n",(0,n.jsx)(i.li,{children:"Download and cache files from the internet for faster access."}),"\n",(0,n.jsx)(i.li,{children:"Accept and store file uploads."}),"\n",(0,n.jsx)(i.li,{children:"Serve files from both cached downloads and uploaded sources."}),"\n"]}),"\n",(0,n.jsx)(i.h2,{id:"use-cases",children:"Use cases"}),"\n",(0,n.jsx)(i.h3,{id:"internal-file-transfer",children:"Internal file transfer"}),"\n",(0,n.jsx)(i.p,{children:'Upload a file from a local node and retrieve it later using the "serve upload" endpoint.'}),"\n",(0,n.jsx)(i.h3,{id:"application-file-uploads",children:"Application file uploads"}),"\n",(0,n.jsx)(i.p,{children:"Enable users or services to upload files, such as profile pictures or media attachments."}),"\n",(0,n.jsx)(i.h3,{id:"component-dependencies",children:"Component dependencies"}),"\n",(0,n.jsx)(i.p,{children:"Download external files (e.g., AI models) via URL and provide them to system components as needed."})]})}function p(e={}){const{wrapper:i}={...(0,o.R)(),...e.components};return i?(0,n.jsx)(i,{...e,children:(0,n.jsx)(c,{...e})}):c(e)}}}]);