"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[3177],{76625:(e,t,s)=>{s.r(t),s.d(t,{assets:()=>m,contentTitle:()=>y,default:()=>k,frontMatter:()=>f,metadata:()=>i,toc:()=>S});const i=JSON.parse('{"id":"1backend/save-grants","title":"Save Grants","description":"Save grants.","source":"@site/docs/1backend/save-grants.api.mdx","sourceDirName":"1backend","slug":"/1backend/save-grants","permalink":"/docs/1backend/save-grants","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"save-grants","title":"Save Grants","description":"Save grants.","sidebar_label":"Save Grants","hide_title":true,"hide_table_of_contents":true,"api":"eJzVVk2P2zYQ/SvEnLXSfjRAq1N3iyDYNmgW6+zJayBjaiwxS5FckrLjGvrvxVBS7NhukV4C9GLQw+HwzZuZR+2goiC9clFZAyXMcE2i9mhiyJ/Ns3mXlqKilTIkNo2SjQi6q4NATwJDULWhSgRHUq2UFI58q0JQ1oRM2DV5ryplahEb4iDY6SikNStVdx75znTLI712ylNIbp+6QP4irGWZcJTSE0b6dBA6hwysoyHAfQUlBFzTABUy8PTaUYh3ttpCuQNpTSQTeYnOaSXTseJz4IR3EGRDLfLKeQ4aFQX+N5DAKxWpDacOquLfuHXEAKJXpoY+gz3M+/MOiT7eOWI+sbppbGI23U5VIuThMGK2h3MSeTSg97iFfm+wy88k4/d4sOlMQ4xd8DgQO1KsPFVQRt9Rz4bgrAkDM9eXl6f5pRiCC1WJ0ElJIaw6rbeQfX+FTtD2Gfx07rJ7s0atKvH77MOf/+WCbytM3lt/hulzvCUkV6dIngx2sbFe/UXVD0Py5jwnkbxBLWbk1+TF2xTzx0DitifZeRW3UM53cEfoyd92sYFyvugXGUTkoZjDUyAvZmsJiwxaio3l8XYdN51Ddodi0oeiniY+pIxCCt15zU6FthJ1Y0Ms3/x8fXMFfMmEYcY5DWkcIjlm7OPWkXgeXZ5BrKzWdkOVWG4FiuBQkkBTiWhfyAiUw0yIlbdtGtwpF/He1soIMpWzykSWL8XxG8KKuAQGWybsduyUVADYT6tTf9AwrlyUx728vf2CrdN0KFfzQZimchzr0d4+ytB8siz6Rc+4VnYSTZSpKahFpZPGagq/BmXqTmP01uTStgfgH+7FrHPOeq7VUIUmRlcWxdUS5QuZig8UcCIxt6IzaqWY18FPVLQmbV1LJgqnMa6sb8XKetEq6S1XW0kKF0sMVInbe3HQt4HZ1UqSCYmXCd27h/difZNffoMtlEWx2Wzy2nS59XUxngsF1k5f3OSXeRNbnYSTfBs+rGbD1fvUwgbrmnyubJFcCi6bilwTuLob0oEMuDuHZC/zm/zywsv8+pf0WtgQWzQHSA8EF46YOnjL/j8P9djEkb7EwmlUhtNOJdiNEz2HKRBkUx8vMuDZ5c3djuv85HXfs/m1I88isshgjV7hkrmec++O48Qj8EJbKOG3ga0LnmMuAupumKcjleuz6cStlOTiv/oe6tLD00fIYDl+aLS24iMeN/xC4gZKSB8qqS3ZIdl2oNHUHdbsO4TkycZRgfYzz4iyacFJTVtmewDwWDOGPPiXszp7ZLcbFKXvv/oPW/944qtQDd5cw0Xf938DdjeR3A==","sidebar_class_name":"put api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"List Grants","permalink":"/docs/1backend/list-grants"},"next":{"title":"Login","permalink":"/docs/1backend/login"}}');var n=s(74848),a=s(28453),r=s(53746),o=s.n(r),d=s(56518),c=s.n(d),p=s(99972),l=s.n(p),u=s(25342),b=s.n(u),g=(s(44215),s(82223),s(24861));const f={id:"save-grants",title:"Save Grants",description:"Save grants.",sidebar_label:"Save Grants",hide_title:!0,hide_table_of_contents:!0,api:"eJzVVk2P2zYQ/SvEnLXSfjRAq1N3iyDYNmgW6+zJayBjaiwxS5FckrLjGvrvxVBS7NhukV4C9GLQw+HwzZuZR+2goiC9clFZAyXMcE2i9mhiyJ/Ns3mXlqKilTIkNo2SjQi6q4NATwJDULWhSgRHUq2UFI58q0JQ1oRM2DV5ryplahEb4iDY6SikNStVdx75znTLI712ylNIbp+6QP4irGWZcJTSE0b6dBA6hwysoyHAfQUlBFzTABUy8PTaUYh3ttpCuQNpTSQTeYnOaSXTseJz4IR3EGRDLfLKeQ4aFQX+N5DAKxWpDacOquLfuHXEAKJXpoY+gz3M+/MOiT7eOWI+sbppbGI23U5VIuThMGK2h3MSeTSg97iFfm+wy88k4/d4sOlMQ4xd8DgQO1KsPFVQRt9Rz4bgrAkDM9eXl6f5pRiCC1WJ0ElJIaw6rbeQfX+FTtD2Gfx07rJ7s0atKvH77MOf/+WCbytM3lt/hulzvCUkV6dIngx2sbFe/UXVD0Py5jwnkbxBLWbk1+TF2xTzx0DitifZeRW3UM53cEfoyd92sYFyvugXGUTkoZjDUyAvZmsJiwxaio3l8XYdN51Ddodi0oeiniY+pIxCCt15zU6FthJ1Y0Ms3/x8fXMFfMmEYcY5DWkcIjlm7OPWkXgeXZ5BrKzWdkOVWG4FiuBQkkBTiWhfyAiUw0yIlbdtGtwpF/He1soIMpWzykSWL8XxG8KKuAQGWybsduyUVADYT6tTf9AwrlyUx728vf2CrdN0KFfzQZimchzr0d4+ytB8siz6Rc+4VnYSTZSpKahFpZPGagq/BmXqTmP01uTStgfgH+7FrHPOeq7VUIUmRlcWxdUS5QuZig8UcCIxt6IzaqWY18FPVLQmbV1LJgqnMa6sb8XKetEq6S1XW0kKF0sMVInbe3HQt4HZ1UqSCYmXCd27h/difZNffoMtlEWx2Wzy2nS59XUxngsF1k5f3OSXeRNbnYSTfBs+rGbD1fvUwgbrmnyubJFcCi6bilwTuLob0oEMuDuHZC/zm/zywsv8+pf0WtgQWzQHSA8EF46YOnjL/j8P9djEkb7EwmlUhtNOJdiNEz2HKRBkUx8vMuDZ5c3djuv85HXfs/m1I88isshgjV7hkrmec++O48Qj8EJbKOG3ga0LnmMuAupumKcjleuz6cStlOTiv/oe6tLD00fIYDl+aLS24iMeN/xC4gZKSB8qqS3ZIdl2oNHUHdbsO4TkycZRgfYzz4iyacFJTVtmewDwWDOGPPiXszp7ZLcbFKXvv/oPW/944qtQDd5cw0Xf938DdjeR3A==",sidebar_class_name:"put api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},y=void 0,m={},S=[];function h(e){const t={code:"code",p:"p",...(0,a.R)(),...e.components};return(0,n.jsxs)(n.Fragment,{children:[(0,n.jsx)(g.default,{as:"h1",className:"openapi__heading",children:"Save Grants"}),"\n",(0,n.jsx)(o(),{method:"put",path:"/user-svc/grants",context:"endpoint"}),"\n",(0,n.jsx)(t.p,{children:"Save grants."}),"\n",(0,n.jsx)(t.p,{children:"Grants define which slugs are assigned specific permissions, overriding the default configuration."}),"\n",(0,n.jsxs)(t.p,{children:["Requires the ",(0,n.jsx)(t.code,{children:"user-svc:grant:create"})," permission."]}),"\n",(0,n.jsx)(g.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,n.jsx)(c(),{parameters:void 0}),"\n",(0,n.jsx)(l(),{title:"Body",body:{content:{"application/json":{schema:{properties:{grants:{items:{properties:{id:{type:"string"},permissionId:{type:"string"},slugs:{description:"Slugs who are granted the PermissionId",items:{type:"string"},type:"array"}},type:"object"},type:"array"}},type:"object"}}},description:"Save Grants Request",required:!0}}),"\n",(0,n.jsx)(b(),{id:void 0,label:void 0,responses:{200:{description:"Grant saved successfully",content:{"application/json":{schema:{type:"object"}}}},400:{description:"Invalid JSON",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}}}})]})}function k(e={}){const{wrapper:t}={...(0,a.R)(),...e.components};return t?(0,n.jsx)(t,{...e,children:(0,n.jsx)(h,{...e})}):h(e)}}}]);