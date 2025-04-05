"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[4010],{16134:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>C,contentTitle:()=>y,default:()=>D,frontMatter:()=>m,metadata:()=>o,toc:()=>M});const o=JSON.parse('{"id":"1backend/get-host","title":"Get Container Host","description":"Retrieve information about the Container host","source":"@site/docs/1backend/get-host.api.mdx","sourceDirName":"1backend","slug":"/1backend/get-host","permalink":"/docs/1backend/get-host","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"get-host","title":"Get Container Host","description":"Retrieve information about the Container host","sidebar_label":"Get Container Host","hide_title":true,"hide_table_of_contents":true,"api":"eJy9VMFu2zAM/RWBZ9dOlg0DfFo3DF3RAS2W9pTlwMisrVaWVElOlhn694G206Zdd9gOuySG9EQ+Pj6yh4qC9MpFZQ2U8I2iV7Qlocyt9S3yscCN7aKIDYlP1kRUhrxobIiQgXXkB9B5BSXUFL+M556CsyZQgLKHN7MZ/z3PdHkBGUhrIpnIt+icVnKIVdwFhvQQZEMt8pfznCmqMeCQvOwh7h1BCSF6ZWpIifM+dMpTBeVqRK2zA8pu7khGSIlxb2fz3yndGOxiY736SdW/kyPvrX+d3WtM3r0mzrmJ5A1qsSS/JS8+DzH/D6WUQSDZeRX3UK56+EjoyZ92sYFytU4sKNaBBX5yw3IrWemWYmMnI0AGDvkNFPKAOwlbWUzOCUNlYUjRec24QluJmu/L+XyxeA+c7MBlybWN5Rwzeqnc9d6R+D5BvoO4tVrbHVVisxcogkNJAk0lor0nI1COdhG33raDw2/CWI74amtlBJnKWWViDhkojt8QVsStMNiycKeTY4ZGwKOe6NQF7QeFeZKY56CCHHpHLSquOKCm8CEoU3cao7cml7Y9in11Lpadc9azYKNITYyuLIr5BuU9mYofFJCyFyqcnp8YjGpLolXSW9ZaSQrCaYw811yOVpJMIOZzyHd29VVsF/nsWbZQFsVut8tr0+XW18X0LhRYO32yyGd5E1vNHCL5NlzeLsdsT2TDDuuafK5sMUAK1klFzZD5x7EQyIDtMNKf5Yt8duJlvnjLcZ0NsUVzxPSM4tEumnbOMwn6p1n566U29TDSj1g4jcowiUGQfvL0Cp55GrLHZTOuphX0/QYD3XidEh8/dOR5nNYZbNEr3HDxq3XKDobiMbinPXdBSnJMY4u6G730YtLT8aidfb6GDHAahif7cbDs8MHRD1dmfxT7pX1HCvybsj886fvR3Ck94serP754nJkRzXquU0q/ABB5ORc=","sidebar_class_name":"get api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Get Container Daemon Information","permalink":"/docs/1backend/container-daemon-info"},"next":{"title":"Build an Image","permalink":"/docs/1backend/build-image"}}');var i=n(74848),s=n(28453),a=n(53746),r=n.n(a),d=n(56518),c=n.n(d),l=n(99972),p=n.n(l),h=n(25342),u=n.n(h),b=(n(44215),n(82223),n(24861));const m={id:"get-host",title:"Get Container Host",description:"Retrieve information about the Container host",sidebar_label:"Get Container Host",hide_title:!0,hide_table_of_contents:!0,api:"eJy9VMFu2zAM/RWBZ9dOlg0DfFo3DF3RAS2W9pTlwMisrVaWVElOlhn694G206Zdd9gOuySG9EQ+Pj6yh4qC9MpFZQ2U8I2iV7Qlocyt9S3yscCN7aKIDYlP1kRUhrxobIiQgXXkB9B5BSXUFL+M556CsyZQgLKHN7MZ/z3PdHkBGUhrIpnIt+icVnKIVdwFhvQQZEMt8pfznCmqMeCQvOwh7h1BCSF6ZWpIifM+dMpTBeVqRK2zA8pu7khGSIlxb2fz3yndGOxiY736SdW/kyPvrX+d3WtM3r0mzrmJ5A1qsSS/JS8+DzH/D6WUQSDZeRX3UK56+EjoyZ92sYFytU4sKNaBBX5yw3IrWemWYmMnI0AGDvkNFPKAOwlbWUzOCUNlYUjRec24QluJmu/L+XyxeA+c7MBlybWN5Rwzeqnc9d6R+D5BvoO4tVrbHVVisxcogkNJAk0lor0nI1COdhG33raDw2/CWI74amtlBJnKWWViDhkojt8QVsStMNiycKeTY4ZGwKOe6NQF7QeFeZKY56CCHHpHLSquOKCm8CEoU3cao7cml7Y9in11Lpadc9azYKNITYyuLIr5BuU9mYofFJCyFyqcnp8YjGpLolXSW9ZaSQrCaYw811yOVpJMIOZzyHd29VVsF/nsWbZQFsVut8tr0+XW18X0LhRYO32yyGd5E1vNHCL5NlzeLsdsT2TDDuuafK5sMUAK1klFzZD5x7EQyIDtMNKf5Yt8duJlvnjLcZ0NsUVzxPSM4tEumnbOMwn6p1n566U29TDSj1g4jcowiUGQfvL0Cp55GrLHZTOuphX0/QYD3XidEh8/dOR5nNYZbNEr3HDxq3XKDobiMbinPXdBSnJMY4u6G730YtLT8aidfb6GDHAahif7cbDs8MHRD1dmfxT7pX1HCvybsj886fvR3Ck94serP754nJkRzXquU0q/ABB5ORc=",sidebar_class_name:"get api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},y=void 0,C={},M=[];function f(e){const t={p:"p",...(0,s.R)(),...e.components};return(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(b.default,{as:"h1",className:"openapi__heading",children:"Get Container Host"}),"\n",(0,i.jsx)(r(),{method:"get",path:"/container-svc/host",context:"endpoint"}),"\n",(0,i.jsx)(t.p,{children:"Retrieve information about the Container host"}),"\n",(0,i.jsx)(c(),{parameters:void 0}),"\n",(0,i.jsx)(p(),{title:"Body",body:void 0}),"\n",(0,i.jsx)(u(),{id:void 0,label:void 0,responses:{200:{description:"OK",content:{"application/json":{schema:{properties:{host:{type:"string"}},required:["host"],type:"object"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}}}})]})}function D(e={}){const{wrapper:t}={...(0,s.R)(),...e.components};return t?(0,i.jsx)(t,{...e,children:(0,i.jsx)(f,{...e})}):f(e)}}}]);