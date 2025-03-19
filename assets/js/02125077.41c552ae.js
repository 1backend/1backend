"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[465],{97900:(t,e,a)=>{a.r(e),a.d(e,{assets:()=>b,contentTitle:()=>f,default:()=>x,frontMatter:()=>h,metadata:()=>s,toc:()=>M});const s=JSON.parse('{"id":"1backend/start-model","title":"Start a Model","description":"Starts a model by ID","source":"@site/docs/1backend/start-model.api.mdx","sourceDirName":"1backend","slug":"/1backend/start-model","permalink":"/docs/1backend/start-model","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"start-model","title":"Start a Model","description":"Starts a model by ID","sidebar_label":"Start a Model","hide_title":true,"hide_table_of_contents":true,"api":"eJy9VVFv2zYQ/ivEPW0AIznJCmx6mosWg9euCebmKfXDmTpbbCiSJSm7nsD/Ppxkx07iDNhLX2xK/Hj3fcfvTj3UFFXQPmlnoYJ5wpCiQNG6moxY7sTsHUhwngIyZFZDBZFBfzEAJHgM2FKiEKG675+FG0BjCM3PHlMDEiy2BBUMOWY1SAj0rdOBaqhS6EhCVA21CFUPaedpSBm0XUPOCwZH72ykyPtXkwn/PU178wEkKGcT2cS76L3RahBQfo0M6V+mcMuvpBLknLOEX86FndkNGl2LP+c3n/5PAh+4fkmPjCkEF85Jk68wuXzJ5M5ilxoX9D9U/zAmb87XJFGwaMScwoaCeD/E/DGUsoRIqgs67QbvvSUMFKZdaqC6X7BVEq7ZlnsfzjcKFhJaSo1jH/suDQZmPJSDGy/iRo2rst/bM5eD34GTscTR510wfKg0TqFpXEzVm1+vri+Bsx5IzVnkqOuU2vMSft55El/2kC8gVs4Yt6Wauw9F9KhIoK1Fcg9kBaqxVcQquFakhsRdpMDaxEe31laQrb3TNhWHpmsIawrHtpvurTPcCDwWFr3+QLuh1NquHPPkS0Q1XCK1qFlxREPx96jtujOYgrOFcu1J7NuZmHfeu6FgY5GalHxVlpdLVA9kaz5QQpbPqjAVndUrzbJHnKhpQ8b5lmwS3mBaudCKlQui1So4vgytKF4sMVItpjNx4rPI4o1WZCMx+wO7P24/is11MXnCLVZlud1ui7XtChfW5f5cLHHtzcV1MSma1BpmnCi08WY1H1MfpcUtrtcUCu3KAVJyVXUyDLl8O8oBCWyeUeykuC4mF0EVV79xXO9iatGeMB3msEBxmLJPatUfu+u1gb2/00TfU+kNastpBsn93vD38Gh4kOMaJFTHmTy6fiGB3c34vudS3wWTM7/+1lHgvltI2GDQuGS5/AXQkdc1VCs0kf6D+09/78f+z+LkQ3GW+sGjdsd1RNPxE0h4oN3JhyQvsjzYnamMu1OlyKeTcy8GUj6dCrd3n0EC7lv12BwcTB4WHP0sp+fNNVLg3yxfOdL3Y+vl/Igft1498djRI5pLtMg5/wtIcad/","sidebar_class_name":"put api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Make a Model Default","permalink":"/docs/1backend/make-default"},"next":{"title":"Get Model Status","permalink":"/docs/1backend/get-model-status"}}');var n=a(74848),d=a(28453),o=a(53746),i=a.n(o),r=a(56518),l=a.n(r),c=a(99972),p=a.n(c),m=a(25342),u=a.n(m),k=(a(44215),a(82223),a(24861));const h={id:"start-model",title:"Start a Model",description:"Starts a model by ID",sidebar_label:"Start a Model",hide_title:!0,hide_table_of_contents:!0,api:"eJy9VVFv2zYQ/ivEPW0AIznJCmx6mosWg9euCebmKfXDmTpbbCiSJSm7nsD/Ppxkx07iDNhLX2xK/Hj3fcfvTj3UFFXQPmlnoYJ5wpCiQNG6moxY7sTsHUhwngIyZFZDBZFBfzEAJHgM2FKiEKG675+FG0BjCM3PHlMDEiy2BBUMOWY1SAj0rdOBaqhS6EhCVA21CFUPaedpSBm0XUPOCwZH72ykyPtXkwn/PU178wEkKGcT2cS76L3RahBQfo0M6V+mcMuvpBLknLOEX86FndkNGl2LP+c3n/5PAh+4fkmPjCkEF85Jk68wuXzJ5M5ilxoX9D9U/zAmb87XJFGwaMScwoaCeD/E/DGUsoRIqgs67QbvvSUMFKZdaqC6X7BVEq7ZlnsfzjcKFhJaSo1jH/suDQZmPJSDGy/iRo2rst/bM5eD34GTscTR510wfKg0TqFpXEzVm1+vri+Bsx5IzVnkqOuU2vMSft55El/2kC8gVs4Yt6Wauw9F9KhIoK1Fcg9kBaqxVcQquFakhsRdpMDaxEe31laQrb3TNhWHpmsIawrHtpvurTPcCDwWFr3+QLuh1NquHPPkS0Q1XCK1qFlxREPx96jtujOYgrOFcu1J7NuZmHfeu6FgY5GalHxVlpdLVA9kaz5QQpbPqjAVndUrzbJHnKhpQ8b5lmwS3mBaudCKlQui1So4vgytKF4sMVItpjNx4rPI4o1WZCMx+wO7P24/is11MXnCLVZlud1ui7XtChfW5f5cLHHtzcV1MSma1BpmnCi08WY1H1MfpcUtrtcUCu3KAVJyVXUyDLl8O8oBCWyeUeykuC4mF0EVV79xXO9iatGeMB3msEBxmLJPatUfu+u1gb2/00TfU+kNastpBsn93vD38Gh4kOMaJFTHmTy6fiGB3c34vudS3wWTM7/+1lHgvltI2GDQuGS5/AXQkdc1VCs0kf6D+09/78f+z+LkQ3GW+sGjdsd1RNPxE0h4oN3JhyQvsjzYnamMu1OlyKeTcy8GUj6dCrd3n0EC7lv12BwcTB4WHP0sp+fNNVLg3yxfOdL3Y+vl/Igft1498djRI5pLtMg5/wtIcad/",sidebar_class_name:"put api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},f=void 0,b={},M=[];function C(t){const e={p:"p",...(0,d.R)(),...t.components};return(0,n.jsxs)(n.Fragment,{children:[(0,n.jsx)(k.default,{as:"h1",className:"openapi__heading",children:"Start a Model"}),"\n",(0,n.jsx)(i(),{method:"put",path:"/model-svc/model/{modelId}/start",context:"endpoint"}),"\n",(0,n.jsx)(e.p,{children:"Starts a model by ID"}),"\n",(0,n.jsx)(k.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,n.jsx)(l(),{parameters:[{description:"Model ID",in:"path",name:"modelId",required:!0,schema:{type:"string"}}]}),"\n",(0,n.jsx)(p(),{title:"Body",body:void 0}),"\n",(0,n.jsx)(u(),{id:void 0,label:void 0,responses:{200:{description:"OK",content:{"application/json":{schema:{type:"object"}}}},400:{description:"Invalid JSON",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}}}})]})}function x(t={}){const{wrapper:e}={...(0,d.R)(),...t.components};return e?(0,n.jsx)(e,{...t,children:(0,n.jsx)(C,{...t})}):C(t)}}}]);