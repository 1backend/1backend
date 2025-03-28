"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[6571],{32109:(e,t,i)=>{i.r(t),i.d(t,{assets:()=>g,contentTitle:()=>f,default:()=>b,frontMatter:()=>h,metadata:()=>s,toc:()=>x});const s=JSON.parse('{"id":"1backend/is-secure","title":"Check Security Status","description":"Returns true if the encryption key is sufficiently secure.","source":"@site/docs/1backend/is-secure.api.mdx","sourceDirName":"1backend","slug":"/1backend/is-secure","permalink":"/docs/1backend/is-secure","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"is-secure","title":"Check Security Status","description":"Returns true if the encryption key is sufficiently secure.","sidebar_label":"Check Security Status","hide_title":true,"hide_table_of_contents":true,"api":"eJylVE2P2zYQ/SvEnGXJGydAoVN3i0WwaIAEq2wvjg9jaiwxpkguSdlVBf33YiT5I25yKPZiC+KbeW8e36iHkoL0ykVlDeTwTLH1JojoWxJqJ2JNgoz03QgQe+qECiK0u52SikzUnQgkW08pJGAdeWTcUwk5qFCMJ5CAp+CsCRQg7+Hdcsl/P/I+ThziL9QtiecZDwlIayKZyBXonFZyJMi+By7rIciaGuQn55k+qonkTJ73EDtHkMPWWk1oYBhY0GurPJWQry/QTXKC2u13khGGgbHvf6b3AUvxTK8thfh/RM79Q/TKVOf+d//t/2KwjbX16h8q30rw4WcDPJlI3qAWBfkDefHovfVvYxoSGLOgYgf5uocHQk/+vo015OvNwPZiFdjygqSnKIqDZNMbirXlxFTEZjrkAsjCCFqEg8xUWIRTlsKoN4wMrdeMzLSVqGsbYv7ht3erO2Cuk5SCFU+huBZ068fXzpH4NkO+gdhZre2RSrHtBIrgUJJAU4po92QEyik/YudtMy7JSyDPA4lPtlJGkCmdVSbyWijuXxOWxAYbbNi3+/l6R3vhnDx06k/qxowqs7Osk28E5Xgj1KDiiQNqCr8HZapWY/TWpNI2V72/PImidc569nMyqY7R5Vl2t0W5J1NyQQZDcuPC/dPCYFQHEo2S3rLXSlIQTmPcWd/wOFpJ4t3M+zPfxy+fxGGVLn9gC3mWHY/HtDJtan2VzXUhw8rpxSpdpnVsNGuI5JvweVdMbBex4YhVRT5VNhshGfukombI3cM0CCTAcZjkL9NVulx4ma6W3NfZEBs0V0r/qEnuRTFHQxQRYxvgxoX+sgRv+x7Odxrp75g5jcqwqNGgfk75Gi4p56Scc75JgPPMiL7fYqAXr4eBX7+25Hm/Ngkc0Cvcsh3rzZCcIsaLsaeO70VKchyBA39Vx3TdbPRwvX4fH79CAjivxyWQ3Cw5PXD305HprnrfBnqSwL9D8ouSvp/iPgxn/HT0y4rzFk1odnQzDMO/RJFTfg==","sidebar_class_name":"get api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Encrypt a Value","permalink":"/docs/1backend/encrypt-value"},"next":{"title":"Remove Secrets","permalink":"/docs/1backend/remove-secrets"}}');var n=i(74848),r=i(28453),c=i(53746),o=i.n(c),a=i(56518),d=i.n(a),u=i(99972),p=i.n(u),l=i(25342),k=i.n(l),y=(i(44215),i(82223),i(24861));const h={id:"is-secure",title:"Check Security Status",description:"Returns true if the encryption key is sufficiently secure.",sidebar_label:"Check Security Status",hide_title:!0,hide_table_of_contents:!0,api:"eJylVE2P2zYQ/SvEnGXJGydAoVN3i0WwaIAEq2wvjg9jaiwxpkguSdlVBf33YiT5I25yKPZiC+KbeW8e36iHkoL0ykVlDeTwTLH1JojoWxJqJ2JNgoz03QgQe+qECiK0u52SikzUnQgkW08pJGAdeWTcUwk5qFCMJ5CAp+CsCRQg7+Hdcsl/P/I+ThziL9QtiecZDwlIayKZyBXonFZyJMi+By7rIciaGuQn55k+qonkTJ73EDtHkMPWWk1oYBhY0GurPJWQry/QTXKC2u13khGGgbHvf6b3AUvxTK8thfh/RM79Q/TKVOf+d//t/2KwjbX16h8q30rw4WcDPJlI3qAWBfkDefHovfVvYxoSGLOgYgf5uocHQk/+vo015OvNwPZiFdjygqSnKIqDZNMbirXlxFTEZjrkAsjCCFqEg8xUWIRTlsKoN4wMrdeMzLSVqGsbYv7ht3erO2Cuk5SCFU+huBZ068fXzpH4NkO+gdhZre2RSrHtBIrgUJJAU4po92QEyik/YudtMy7JSyDPA4lPtlJGkCmdVSbyWijuXxOWxAYbbNi3+/l6R3vhnDx06k/qxowqs7Osk28E5Xgj1KDiiQNqCr8HZapWY/TWpNI2V72/PImidc569nMyqY7R5Vl2t0W5J1NyQQZDcuPC/dPCYFQHEo2S3rLXSlIQTmPcWd/wOFpJ4t3M+zPfxy+fxGGVLn9gC3mWHY/HtDJtan2VzXUhw8rpxSpdpnVsNGuI5JvweVdMbBex4YhVRT5VNhshGfukombI3cM0CCTAcZjkL9NVulx4ma6W3NfZEBs0V0r/qEnuRTFHQxQRYxvgxoX+sgRv+x7Odxrp75g5jcqwqNGgfk75Gi4p56Scc75JgPPMiL7fYqAXr4eBX7+25Hm/Ngkc0Cvcsh3rzZCcIsaLsaeO70VKchyBA39Vx3TdbPRwvX4fH79CAjivxyWQ3Cw5PXD305HprnrfBnqSwL9D8ouSvp/iPgxn/HT0y4rzFk1odnQzDMO/RJFTfg==",sidebar_class_name:"get api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},f=void 0,g={},x=[];function C(e){const t={p:"p",...(0,r.R)(),...e.components};return(0,n.jsxs)(n.Fragment,{children:[(0,n.jsx)(y.default,{as:"h1",className:"openapi__heading",children:"Check Security Status"}),"\n",(0,n.jsx)(o(),{method:"get",path:"/secret-svc/is-secure",context:"endpoint"}),"\n",(0,n.jsx)(t.p,{children:"Returns true if the encryption key is sufficiently secure."}),"\n",(0,n.jsx)(d(),{parameters:void 0}),"\n",(0,n.jsx)(p(),{title:"Body",body:void 0}),"\n",(0,n.jsx)(k(),{id:void 0,label:void 0,responses:{200:{description:"Encrypt Value Response",content:{"application/json":{schema:{properties:{isSecure:{type:"boolean"}},required:["isSecure"],type:"object"}}}},400:{description:"Bad Request",content:{"application/json":{schema:{type:"string"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{type:"string"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{type:"string"}}}}}})]})}function b(e={}){const{wrapper:t}={...(0,r.R)(),...e.components};return t?(0,n.jsx)(t,{...e,children:(0,n.jsx)(C,{...e})}):C(e)}}}]);