"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[3454],{84436:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>k,contentTitle:()=>E,default:()=>V,frontMatter:()=>b,metadata:()=>s,toc:()=>f});const s=JSON.parse('{"id":"1backend/encrypt-value","title":"Encrypt a Value","description":"Encrypt a value and return the encrypted result","source":"@site/docs/1backend/encrypt-value.api.mdx","sourceDirName":"1backend","slug":"/1backend/encrypt-value","permalink":"/docs/1backend/encrypt-value","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"encrypt-value","title":"Encrypt a Value","description":"Encrypt a value and return the encrypted result","sidebar_label":"Encrypt a Value","hide_title":true,"hide_table_of_contents":true,"api":"eJzNVU2P2zgM/SsCz46dbFoU8GlnFoNi0AIdNJ29pDkwMuOoY0sqJSfNGv7vC/pjkk7TxWJ72UsiU6T4+MRHtVBQ0Gx8NM5CDndW88lHheqAVUMKbaGYYsNWxT0pGrZJjKGpIiTgPDFK9H0BOYwOf0owJMD0taEQb11xgrwF7WwkG2WJ3ldG94HZlyC5Wwh6TzXKyrMcGw0F+eqhyCKePEEOIbKxJXTJsNP7mEh1uOozGpAZT9CdDW77hXSETkzXSeirUB+HGsZqDFMBeeSGOjEE72wYEPw2n8vfPx81+EPyf6WiS+DVtTpusbgg4l+Df4FgPH/x4/mPFpu4d2z+ouJXE7y+VsC9jcQWK7UiPhCrO2bHv5apSyCQbtjEE+TrFm4JmfimiXvI15tuk0DEMkC+hhVppqhWBw2bBGqKeydi8a5n06NEQBZ6r1k46GyUEUgGgRv6BA1X4pdVTmO1dyHmi8Vy+QYk1YRkJYCHPrjE85KOTydP6vPo8hnUzlWVO1KhtieFKnjUg/ajeyKrUA+Nr3bs6n4QPAZiqUe9d6WximzhnbExhQSMnL8nLEj4tVgLbTfj7fbswrkTvXlHQysK4x/P4+LuG9a+oouen6g/t/p6Mm06Sbtz04xB3V8o1WiEsYAVhd+DsWVTYWRnU+3qC2wP92rVeO9YCB9I3sfo8yxbbFE/kS0kIIMfBsXN/cxiNAdStdHs5K6MpqB8hXHnuBY6KqNJJJ+3z/nePrxXh2U6/y5byLPseDympW1Sx2U2xoUMS1/Nluk83ce66lVMXIcPu9WQ7Qw2HLEsiVPjst4lE55NFBJhcTsUIuwRhwH+PF2m8xnrdPlKzpV2rNFeID2/BtNA/67+i4H+Hx6OsQUifYuZr9BYwdDz0Y6SWMNZEpBMb4tISJpf9tt2i4Eeueo6MX9tiEWLG+kRNriV2tfSHWM/ioqe6AQ5/DEgn4kQppbqG/LFDOiSKeJGa+ol+XPfS3E/fFh9ggS249NXu0JiGI/ykOARcugfT4nu1drbWqjQlg2W4jucKdrAUcNn1QikZFpIVdOWPV0gfKm6oRD5lbKuhrTtoMmue/Yftn4a8Sz1wVvucdN13d92HOAH","sidebar_class_name":"post api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Decrypt a Value","permalink":"/docs/1backend/decrypt-value"},"next":{"title":"Check Security Status","permalink":"/docs/1backend/is-secure"}}');var r=n(74848),i=n(28453),a=n(53746),p=n.n(a),c=n(56518),d=n.n(c),o=n(99972),l=n.n(o),u=n(25342),y=n.n(u),h=(n(44215),n(82223),n(24861));const b={id:"encrypt-value",title:"Encrypt a Value",description:"Encrypt a value and return the encrypted result",sidebar_label:"Encrypt a Value",hide_title:!0,hide_table_of_contents:!0,api:"eJzNVU2P2zgM/SsCz46dbFoU8GlnFoNi0AIdNJ29pDkwMuOoY0sqJSfNGv7vC/pjkk7TxWJ72UsiU6T4+MRHtVBQ0Gx8NM5CDndW88lHheqAVUMKbaGYYsNWxT0pGrZJjKGpIiTgPDFK9H0BOYwOf0owJMD0taEQb11xgrwF7WwkG2WJ3ldG94HZlyC5Wwh6TzXKyrMcGw0F+eqhyCKePEEOIbKxJXTJsNP7mEh1uOozGpAZT9CdDW77hXSETkzXSeirUB+HGsZqDFMBeeSGOjEE72wYEPw2n8vfPx81+EPyf6WiS+DVtTpusbgg4l+Df4FgPH/x4/mPFpu4d2z+ouJXE7y+VsC9jcQWK7UiPhCrO2bHv5apSyCQbtjEE+TrFm4JmfimiXvI15tuk0DEMkC+hhVppqhWBw2bBGqKeydi8a5n06NEQBZ6r1k46GyUEUgGgRv6BA1X4pdVTmO1dyHmi8Vy+QYk1YRkJYCHPrjE85KOTydP6vPo8hnUzlWVO1KhtieFKnjUg/ajeyKrUA+Nr3bs6n4QPAZiqUe9d6WximzhnbExhQSMnL8nLEj4tVgLbTfj7fbswrkTvXlHQysK4x/P4+LuG9a+oouen6g/t/p6Mm06Sbtz04xB3V8o1WiEsYAVhd+DsWVTYWRnU+3qC2wP92rVeO9YCB9I3sfo8yxbbFE/kS0kIIMfBsXN/cxiNAdStdHs5K6MpqB8hXHnuBY6KqNJJJ+3z/nePrxXh2U6/y5byLPseDympW1Sx2U2xoUMS1/Nluk83ce66lVMXIcPu9WQ7Qw2HLEsiVPjst4lE55NFBJhcTsUIuwRhwH+PF2m8xnrdPlKzpV2rNFeID2/BtNA/67+i4H+Hx6OsQUifYuZr9BYwdDz0Y6SWMNZEpBMb4tISJpf9tt2i4Eeueo6MX9tiEWLG+kRNriV2tfSHWM/ioqe6AQ5/DEgn4kQppbqG/LFDOiSKeJGa+ol+XPfS3E/fFh9ggS249NXu0JiGI/ykOARcugfT4nu1drbWqjQlg2W4jucKdrAUcNn1QikZFpIVdOWPV0gfKm6oRD5lbKuhrTtoMmue/Yftn4a8Sz1wVvucdN13d92HOAH",sidebar_class_name:"post api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},E=void 0,k={},f=[];function m(e){const t={p:"p",...(0,i.R)(),...e.components};return(0,r.jsxs)(r.Fragment,{children:[(0,r.jsx)(h.default,{as:"h1",className:"openapi__heading",children:"Encrypt a Value"}),"\n",(0,r.jsx)(p(),{method:"post",path:"/secret-svc/encrypt",context:"endpoint"}),"\n",(0,r.jsx)(t.p,{children:"Encrypt a value and return the encrypted result"}),"\n",(0,r.jsx)(h.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,r.jsx)(d(),{parameters:void 0}),"\n",(0,r.jsx)(l(),{title:"Body",body:{content:{"application/json":{schema:{properties:{value:{type:"string"},values:{items:{type:"string"},type:"array"}},type:"object"}}},description:"Encrypt Value Request",required:!0}}),"\n",(0,r.jsx)(y(),{id:void 0,label:void 0,responses:{200:{description:"Encrypt Value Response",content:{"application/json":{schema:{properties:{value:{type:"string"},values:{items:{type:"string"},type:"array"}},type:"object"}}}},400:{description:"Bad Request",content:{"application/json":{schema:{type:"string"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{type:"string"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{type:"string"}}}}}})]})}function V(e={}){const{wrapper:t}={...(0,i.R)(),...e.components};return t?(0,r.jsx)(t,{...e,children:(0,r.jsx)(m,{...e})}):m(e)}}}]);