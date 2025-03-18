"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[9299],{5619:(e,n,a)=>{a.r(n),a.d(n,{assets:()=>f,contentTitle:()=>g,default:()=>B,frontMatter:()=>j,metadata:()=>i,toc:()=>h});const i=JSON.parse('{"id":"1backend/image-pullable","title":"Check if Container Image is Pullable","description":"Check if an image exists on in the container registry and is pullable.","source":"@site/docs/1backend/image-pullable.api.mdx","sourceDirName":"1backend","slug":"/1backend/image-pullable","permalink":"/docs/1backend/image-pullable","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"image-pullable","title":"Check if Container Image is Pullable","description":"Check if an image exists on in the container registry and is pullable.","sidebar_label":"Check if Container Image is Pullable","hide_title":true,"hide_table_of_contents":true,"api":"eJy9VU2P2zYQ/SvEnFpAK3l3G6DVqU4QBEaCZlFnTxsfaGosMUuRDEnZcQUC+RH5hfklxVCyLW+cHoogF3+Iw+F7j++NeqjQCydtkEZDCS8aFI9MbhjXTLa8RoafpA+eGc2kZqFBJowOXGp0zGEtfXB7xnXFpGe2U4qvFeaQgbHoODVdVFBCanU3LkMGljveYkDnoXzon2BYpHM1b6lS0hPLQwMZpEdjs7+GZYcfO+mwgjK4DjPwosGWQ9lD2Fsq9sFJXUOMKyr21miPntZvZjP6Oj/67WvIgAiiDrTKrVVSJB7FB08l/eQI64hlkEPDA/vJ4WtjFHINMU6RPpxKV9mh1Kw/oAgQI9X+Nrv+Ftu95l1ojJP/YPX/UaJzxl3S5zKSZ5dUWuiATnPFlui26NjL1PPnQIoZeBSdk2GfrPMcuUM370ID5cOKbjnwmlwFL44+XW4FKd1iaAy5scaQPEh7oDj6+cpvRZHMVfRHj8XCnmzrE9/Bs51TtLtQRnDVGB/KZ7/f3F4DQTggXBLjgeQU51M93+0tsvdjyXtgG6OU2WHF1nvGmbdcYIpYMI+oGReDkdjGmTYl8t4PJNkbU0vNUFfWSB3yQ3wa5BW6U4Dmo4/S9cBRZW7la9wn3aXeGMKZtBHpRrHlkhh7rtD/6aWuO8WDMzoXpp30vluwZWetcaTxIFITgi2L4nrNxSPqijYUELMnKsxZp+VGEu2hjlW4RWVsizowq3jYGNeyjXFsbzrH5gs2cZn/+vlLK4UzdEdSoL9ac49V0m3dSRVYMMwLPkwnJQVqn6J6wP3q7g3b3uazM9S+LIrdbpfXusuNq4txny94bdXVbT7Lm9Aq4hLQtf7tZjmcfiLtd7yu0eXSFKmkIL1loDEB188HopAB2WqQYZbf5rMrJ/KbP6ivNT60XE+QHkf0yeDDxJSeTUbsmbj9KZs/bsSPtgn4KRRWcakJb9KuH8P1AGfhIj/SgZBBOR3i03FISaKNfU/3d+9UjPT4Y4eOAr/KYMudHOYsvTmkp98VlBuuPP4H7V/+Hgfwr+zsBXORxSERek93w1VH/yCDR9yfvYDiKmaHeBGcYX0uBNow2fnNNIzTcfTq5TvIgI+j4RRGapYdflD3i6iehnmAQJ8x+86Wvh+iHuOxflj67o7jBBkFIAYxxn8B5B/VKA==","sidebar_class_name":"get api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Build an Image","permalink":"/docs/1backend/build-image"},"next":{"title":"List Logs","permalink":"/docs/1backend/list-container-logs"}}');var t=a(74848),l=a(28453),s=a(53746),o=a.n(s),r=a(56518),c=a.n(r),d=a(99972),p=a.n(d),u=a(25342),b=a.n(u),m=(a(44215),a(82223),a(24861));const j={id:"image-pullable",title:"Check if Container Image is Pullable",description:"Check if an image exists on in the container registry and is pullable.",sidebar_label:"Check if Container Image is Pullable",hide_title:!0,hide_table_of_contents:!0,api:"eJy9VU2P2zYQ/SvEnFpAK3l3G6DVqU4QBEaCZlFnTxsfaGosMUuRDEnZcQUC+RH5hfklxVCyLW+cHoogF3+Iw+F7j++NeqjQCydtkEZDCS8aFI9MbhjXTLa8RoafpA+eGc2kZqFBJowOXGp0zGEtfXB7xnXFpGe2U4qvFeaQgbHoODVdVFBCanU3LkMGljveYkDnoXzon2BYpHM1b6lS0hPLQwMZpEdjs7+GZYcfO+mwgjK4DjPwosGWQ9lD2Fsq9sFJXUOMKyr21miPntZvZjP6Oj/67WvIgAiiDrTKrVVSJB7FB08l/eQI64hlkEPDA/vJ4WtjFHINMU6RPpxKV9mh1Kw/oAgQI9X+Nrv+Ftu95l1ojJP/YPX/UaJzxl3S5zKSZ5dUWuiATnPFlui26NjL1PPnQIoZeBSdk2GfrPMcuUM370ID5cOKbjnwmlwFL44+XW4FKd1iaAy5scaQPEh7oDj6+cpvRZHMVfRHj8XCnmzrE9/Bs51TtLtQRnDVGB/KZ7/f3F4DQTggXBLjgeQU51M93+0tsvdjyXtgG6OU2WHF1nvGmbdcYIpYMI+oGReDkdjGmTYl8t4PJNkbU0vNUFfWSB3yQ3wa5BW6U4Dmo4/S9cBRZW7la9wn3aXeGMKZtBHpRrHlkhh7rtD/6aWuO8WDMzoXpp30vluwZWetcaTxIFITgi2L4nrNxSPqijYUELMnKsxZp+VGEu2hjlW4RWVsizowq3jYGNeyjXFsbzrH5gs2cZn/+vlLK4UzdEdSoL9ac49V0m3dSRVYMMwLPkwnJQVqn6J6wP3q7g3b3uazM9S+LIrdbpfXusuNq4txny94bdXVbT7Lm9Aq4hLQtf7tZjmcfiLtd7yu0eXSFKmkIL1loDEB188HopAB2WqQYZbf5rMrJ/KbP6ivNT60XE+QHkf0yeDDxJSeTUbsmbj9KZs/bsSPtgn4KRRWcakJb9KuH8P1AGfhIj/SgZBBOR3i03FISaKNfU/3d+9UjPT4Y4eOAr/KYMudHOYsvTmkp98VlBuuPP4H7V/+Hgfwr+zsBXORxSERek93w1VH/yCDR9yfvYDiKmaHeBGcYX0uBNow2fnNNIzTcfTq5TvIgI+j4RRGapYdflD3i6iehnmAQJ8x+86Wvh+iHuOxflj67o7jBBkFIAYxxn8B5B/VKA==",sidebar_class_name:"get api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},g=void 0,f={},h=[];function k(e){const n={p:"p",...(0,l.R)(),...e.components};return(0,t.jsxs)(t.Fragment,{children:[(0,t.jsx)(m.default,{as:"h1",className:"openapi__heading",children:"Check if Container Image is Pullable"}),"\n",(0,t.jsx)(o(),{method:"get",path:"/container-svc/image/{imageName}/pullable",context:"endpoint"}),"\n",(0,t.jsx)(n.p,{children:"Check if an image exists on in the container registry and is pullable."}),"\n",(0,t.jsx)(m.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,t.jsx)(c(),{parameters:[{description:"Image name",in:"path",name:"imageName",required:!0,schema:{type:"string"}}]}),"\n",(0,t.jsx)(p(),{title:"Body",body:void 0}),"\n",(0,t.jsx)(b(),{id:void 0,label:void 0,responses:{200:{description:"OK",content:{"application/json":{schema:{properties:{pullable:{type:"boolean"}},required:["pullable"],type:"object"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}}}})]})}function B(e={}){const{wrapper:n}={...(0,l.R)(),...e.components};return n?(0,t.jsx)(n,{...e,children:(0,t.jsx)(k,{...e})}):k(e)}}}]);