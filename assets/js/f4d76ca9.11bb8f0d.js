"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[6653],{28367:(e,t,s)=>{s.r(t),s.d(t,{assets:()=>f,contentTitle:()=>y,default:()=>q,frontMatter:()=>k,metadata:()=>c,toc:()=>m});const c=JSON.parse('{"id":"1backend/check","title":"Check","description":"Check records a resource access and returns if the access is allowed.","source":"@site/docs/1backend/check.api.mdx","sourceDirName":"1backend","slug":"/1backend/check","permalink":"/docs/1backend/check","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"check","title":"Check","description":"Check records a resource access and returns if the access is allowed.","sidebar_label":"Check","hide_title":true,"hide_table_of_contents":true,"api":"eJy9Vs2O2zYQfhVizlrJu9sArU7dDYJi26C7iLMnx4cxNZaZpUiGpOyohoA8RJ+wT1IMJdnejVsUQZGLLQ1nxt/3zQ+9h4qC9MpFZQ2U8HpD8kl4ktZXQaDwFGzrJQmUkkIQaCrhKbbeBKHWIm4OJyoI1NruqMohA+vII+e8q6AEyVkhA0+fWgrx1lYdlHuQ1kQykR/ROa1kiig+BoayhyA31CA/Oc/5oqLAb2QqZ9UQFztHUEKIXpka+gyUO2tuKG5sdfaoDeTvzh312WSxq48kI/RsOqfXu4HXyFB5qqCMvqWeDcFZEwbkV7MZf51JQZUIbRJy3WrdQfat4ow1OKGzslYTmsTnCG9x8Fyeodln8MM5rHdmi1pV4tf5/e/fjpG8t/4/Cp6QXH6N5NFgGzfWqz+o+m5IXp3XJJI3qMWc/Ja8eJNyfh9IfQaBZOtV7KBc7OGW0JO/aeMGysWy59JiHbjaD1Yr2Yn5VnLBp3EAZ1PbOuQIKFzyughbWUwzGxKrkNK3XrNXoa1EvbEhlq9+vLq+BP6hCceceQ1UTtG8VO1950h8GF0+gFjboR3FqhMogkPeOaYS0T6RESiHvhVrb5u0dR4DeWYj3tpaGTHtBN49ivNvCCviMhhsWLSbsVtSEeCgJTr1G3VJXS7Mu+OCevMZG6fp+cKZqjHsmePbQc+DZdoqp6vJrO209lCmxqAGFUsaUFP4OShTtxqjtyaXtjkB/3An5q1z1nOxhipsYnRlUVyuUD6RqTiggK/2042QtmmsEaObWFsvOtt6cXMnTtoy/PXlz0ZJb7ncSlK4WGGgKpVg1SodRbQiSNTECmslyYSkzYTwl4e3Ynudz57hC2VR7Ha7vDZtbn1djHGhwNrpi+t8lm9ioxl1JN+E+/V8+PUjvbDDuiafK1skl4JLpyLXBS5vB06QAXfoQHiWX+ezCy/zq584L/d3g+YE6euxr5/pdHIX/W9X4NhhkT7HwmlUJl02rM1+nLcFHOeNF0ZCtsyAJ4tP93suwqPXfc/mTy15HvNlBlv0CleswmLZZ1Oz84g+UcccBjYXPGUsD+p26PYXe6jPpogbKcnFf/U93RsP9/P3kMFqvMkbW3GMxx3fgbiDEtKfgNRa7JBse9Bo6hZr9h1y8uDhuCCOI8mQsumBWU1HpjtB+HKkByL8ybTOhuz3w8D3/cF/OPrHiMMeGby5isu+7/8Grs0s6w==","sidebar_class_name":"post api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"List Platforms","permalink":"/docs/1backend/list-platforms"},"next":{"title":"Upsert an Instance","permalink":"/docs/1backend/upsert-instance"}}');var n=s(74848),a=s(28453),i=s(53746),r=s.n(i),o=s(56518),p=s.n(o),d=s(99972),l=s.n(d),h=s(25342),u=s.n(h),b=(s(44215),s(82223),s(24861));const k={id:"check",title:"Check",description:"Check records a resource access and returns if the access is allowed.",sidebar_label:"Check",hide_title:!0,hide_table_of_contents:!0,api:"eJy9Vs2O2zYQfhVizlrJu9sArU7dDYJi26C7iLMnx4cxNZaZpUiGpOyohoA8RJ+wT1IMJdnejVsUQZGLLQ1nxt/3zQ+9h4qC9MpFZQ2U8HpD8kl4ktZXQaDwFGzrJQmUkkIQaCrhKbbeBKHWIm4OJyoI1NruqMohA+vII+e8q6AEyVkhA0+fWgrx1lYdlHuQ1kQykR/ROa1kiig+BoayhyA31CA/Oc/5oqLAb2QqZ9UQFztHUEKIXpka+gyUO2tuKG5sdfaoDeTvzh312WSxq48kI/RsOqfXu4HXyFB5qqCMvqWeDcFZEwbkV7MZf51JQZUIbRJy3WrdQfat4ow1OKGzslYTmsTnCG9x8Fyeodln8MM5rHdmi1pV4tf5/e/fjpG8t/4/Cp6QXH6N5NFgGzfWqz+o+m5IXp3XJJI3qMWc/Ja8eJNyfh9IfQaBZOtV7KBc7OGW0JO/aeMGysWy59JiHbjaD1Yr2Yn5VnLBp3EAZ1PbOuQIKFzyughbWUwzGxKrkNK3XrNXoa1EvbEhlq9+vLq+BP6hCceceQ1UTtG8VO1950h8GF0+gFjboR3FqhMogkPeOaYS0T6RESiHvhVrb5u0dR4DeWYj3tpaGTHtBN49ivNvCCviMhhsWLSbsVtSEeCgJTr1G3VJXS7Mu+OCevMZG6fp+cKZqjHsmePbQc+DZdoqp6vJrO209lCmxqAGFUsaUFP4OShTtxqjtyaXtjkB/3An5q1z1nOxhipsYnRlUVyuUD6RqTiggK/2042QtmmsEaObWFsvOtt6cXMnTtoy/PXlz0ZJb7ncSlK4WGGgKpVg1SodRbQiSNTECmslyYSkzYTwl4e3Ynudz57hC2VR7Ha7vDZtbn1djHGhwNrpi+t8lm9ioxl1JN+E+/V8+PUjvbDDuiafK1skl4JLpyLXBS5vB06QAXfoQHiWX+ezCy/zq584L/d3g+YE6euxr5/pdHIX/W9X4NhhkT7HwmlUJl02rM1+nLcFHOeNF0ZCtsyAJ4tP93suwqPXfc/mTy15HvNlBlv0CleswmLZZ1Oz84g+UcccBjYXPGUsD+p26PYXe6jPpogbKcnFf/U93RsP9/P3kMFqvMkbW3GMxx3fgbiDEtKfgNRa7JBse9Bo6hZr9h1y8uDhuCCOI8mQsumBWU1HpjtB+HKkByL8ybTOhuz3w8D3/cF/OPrHiMMeGby5isu+7/8Grs0s6w==",sidebar_class_name:"post api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},y=void 0,f={},m=[];function C(e){const t={p:"p",...(0,a.R)(),...e.components};return(0,n.jsxs)(n.Fragment,{children:[(0,n.jsx)(b.default,{as:"h1",className:"openapi__heading",children:"Check"}),"\n",(0,n.jsx)(r(),{method:"post",path:"/policy-svc/check",context:"endpoint"}),"\n",(0,n.jsx)(t.p,{children:"Check records a resource access and returns if the access is allowed."}),"\n",(0,n.jsx)(b.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,n.jsx)(p(),{parameters:void 0}),"\n",(0,n.jsx)(l(),{title:"Body",body:{content:{"application/json":{schema:{properties:{endpoint:{type:"string"},ip:{type:"string"},method:{type:"string"},userId:{type:"string"}},type:"object"}}},description:"Check Request",required:!0}}),"\n",(0,n.jsx)(u(),{id:void 0,label:void 0,responses:{200:{description:"Checked successfully",content:{"application/json":{schema:{properties:{allowed:{type:"boolean"}},required:["allowed"],type:"object"}}}},400:{description:"Invalid JSON",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}}}})]})}function q(e={}){const{wrapper:t}={...(0,a.R)(),...e.components};return t?(0,n.jsx)(t,{...e,children:(0,n.jsx)(C,{...e})}):C(e)}}}]);