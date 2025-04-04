"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[4010],{16134:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>m,contentTitle:()=>k,default:()=>S,frontMatter:()=>C,metadata:()=>o,toc:()=>y});const o=JSON.parse('{"id":"1backend/get-host","title":"Get Container Host","description":"Retrieve information about the Container host","source":"@site/docs/1backend/get-host.api.mdx","sourceDirName":"1backend","slug":"/1backend/get-host","permalink":"/docs/1backend/get-host","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"get-host","title":"Get Container Host","description":"Retrieve information about the Container host","sidebar_label":"Get Container Host","hide_title":true,"hide_table_of_contents":true,"api":"eJy9VMFu2zAM/RWBZ9dOmxUYfFo7DF2xAS2W9ZTmwMisrVWWVElOlhn694G206Zde9gOuySG9EQ+Pj6yh4qC9MpFZQ2U8I2iV7Qhocyd9S3yscC17aKIDYmP1kRUhrxobIiQgXXkB9BlBSXUFD+P556CsyZQgLKHk9mM/55nuvoCGUhrIpnIt+icVnKIVfwIDOkhyIZa5C/nOVNUY8AhedlD3DmCEkL0ytSQEud96JSnCsrliFple5Rd/yAZISXGvZsd/0npxmAXG+vVL6r+nRx5b/3r7F5jcvqaOJcmkjeoxYL8hrz4NMT8P5RSBoFk51XcQbns4ZzQkz/rYgPlcpVYUKwDC/zkhsVGstItxcZORoAMHPIbKOQedxQ2spicE4bKwpCi85pxhbYSNd+Xp+9P5sfAyfZcFlzbWM4ho5fKfd85ErcT5BbEndXabqkS651AERxKEmgqEe09GYFytIu487YdHH4TxnLEV1srI8hUzioTc8hAcfyGsCJuhcGWhTubHDM0Ah71RKe+0G5QmCeJeQ4qyKF31KLiigNqCh+CMnWnMXprcmnbg9jXl2LROWc9CzaK1MToyqI4XqO8J1PxgwJS9kKFs8sjg1FtSLRKestaK0lBOI2R55rL0UqSCcR89vkurr+KzTyfPcsWyqLYbrd5bbrc+rqY3oUCa6eP5vksb2KrmUMk34aru8WY7Yls2GJdk8+VLQZIwTqpqBlyfD4WAhmwHUb6s3yez468zOdzjutsiC2aA6YXFA920bRznknQP83KXy+1qYeRfsbCaVSGSQyC9JOnl/DM05A9LptxNS2h79cY6MbrlPj4oSPP47TKYINe4ZqLX65StjcUj8E97bgLUpJjGhvU3eilF5OeDkft4tN3yACnYXiyHwfL9h8cfX9ldgexX9p3pMC/KXvjSd+P5k7pET9evfnicWZGNOu5Sin9Bh5AORo=","sidebar_class_name":"get api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Get Container Daemon Information","permalink":"/docs/1backend/container-daemon-info"},"next":{"title":"Build an Image","permalink":"/docs/1backend/build-image"}}');var i=n(74848),s=n(28453),r=n(53746),a=n.n(r),c=n(56518),d=n.n(c),p=n(99972),h=n.n(p),l=n(25342),b=n.n(l),u=(n(44215),n(82223),n(24861));const C={id:"get-host",title:"Get Container Host",description:"Retrieve information about the Container host",sidebar_label:"Get Container Host",hide_title:!0,hide_table_of_contents:!0,api:"eJy9VMFu2zAM/RWBZ9dOmxUYfFo7DF2xAS2W9ZTmwMisrVWWVElOlhn694G206Zde9gOuySG9EQ+Pj6yh4qC9MpFZQ2U8I2iV7Qhocyd9S3yscC17aKIDYmP1kRUhrxobIiQgXXkB9BlBSXUFD+P556CsyZQgLKHk9mM/55nuvoCGUhrIpnIt+icVnKIVfwIDOkhyIZa5C/nOVNUY8AhedlD3DmCEkL0ytSQEud96JSnCsrliFple5Rd/yAZISXGvZsd/0npxmAXG+vVL6r+nRx5b/3r7F5jcvqaOJcmkjeoxYL8hrz4NMT8P5RSBoFk51XcQbns4ZzQkz/rYgPlcpVYUKwDC/zkhsVGstItxcZORoAMHPIbKOQedxQ2spicE4bKwpCi85pxhbYSNd+Xp+9P5sfAyfZcFlzbWM4ho5fKfd85ErcT5BbEndXabqkS651AERxKEmgqEe09GYFytIu487YdHH4TxnLEV1srI8hUzioTc8hAcfyGsCJuhcGWhTubHDM0Ah71RKe+0G5QmCeJeQ4qyKF31KLiigNqCh+CMnWnMXprcmnbg9jXl2LROWc9CzaK1MToyqI4XqO8J1PxgwJS9kKFs8sjg1FtSLRKestaK0lBOI2R55rL0UqSCcR89vkurr+KzTyfPcsWyqLYbrd5bbrc+rqY3oUCa6eP5vksb2KrmUMk34aru8WY7Yls2GJdk8+VLQZIwTqpqBlyfD4WAhmwHUb6s3yez468zOdzjutsiC2aA6YXFA920bRznknQP83KXy+1qYeRfsbCaVSGSQyC9JOnl/DM05A9LptxNS2h79cY6MbrlPj4oSPP47TKYINe4ZqLX65StjcUj8E97bgLUpJjGhvU3eilF5OeDkft4tN3yACnYXiyHwfL9h8cfX9ldgexX9p3pMC/KXvjSd+P5k7pET9evfnicWZGNOu5Sin9Bh5AORo=",sidebar_class_name:"get api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},k=void 0,m={},y=[];function f(e){const t={p:"p",...(0,s.R)(),...e.components};return(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(u.default,{as:"h1",className:"openapi__heading",children:"Get Container Host"}),"\n",(0,i.jsx)(a(),{method:"get",path:"/container-svc/host",context:"endpoint"}),"\n",(0,i.jsx)(t.p,{children:"Retrieve information about the Container host"}),"\n",(0,i.jsx)(d(),{parameters:void 0}),"\n",(0,i.jsx)(h(),{title:"Body",body:void 0}),"\n",(0,i.jsx)(b(),{id:void 0,label:void 0,responses:{200:{description:"OK",content:{"application/json":{schema:{properties:{host:{type:"string"}},required:["host"],type:"object"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}}}})]})}function S(e={}){const{wrapper:t}={...(0,s.R)(),...e.components};return t?(0,i.jsx)(t,{...e,children:(0,i.jsx)(f,{...e})}):f(e)}}}]);