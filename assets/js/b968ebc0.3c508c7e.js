"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[6842],{56291:(e,t,s)=>{s.r(t),s.d(t,{assets:()=>f,contentTitle:()=>b,default:()=>M,frontMatter:()=>u,metadata:()=>a,toc:()=>y});const a=JSON.parse('{"id":"1backend/get-message","title":"Get Message","description":"Fetch information about a specific chat message by its ID","source":"@site/docs/1backend/get-message.api.mdx","sourceDirName":"1backend","slug":"/1backend/get-message","permalink":"/docs/1backend/get-message","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"get-message","title":"Get Message","description":"Fetch information about a specific chat message by its ID","sidebar_label":"Get Message","hide_title":true,"hide_table_of_contents":true,"api":"eJylVltv2zYU/isEX7YBiuQ0K1D4ZXO3JPPWLsHcoA+JMRxTxxIbimTJIzueof8+HEmO7VgtUPTFl8Nz+c7tI7cyx6iC9qSdlWN5haRKoe3ShQpYJmDhahIgokell1oJVQKJCmOEAsViIzRFMf1dJtJ5DK3NNJdjWSC975RkIj0EqJAwRDm+376I2at1TjRLPFApE2mhQjmWfaxpLhMZ8HOtA+ZyTKHGREZVYgVyvJW08awcKWhbyKaZs3L0zkaMfP5qNOKv4dA5EmgTRayVwhiXtTEbEZCCxhVyXOUsoSX2AN4brdpEs0+R3WwPYPjAZSDdBcUnHSkewFs4ZxCsbJJdWqdGKiAQ5hMaSCuRS21wmsfTXK66A5HjUluMgkoUrCyACFRZoaVOuOtdCTHlihNWcTBUL4AQYMP/dd7lBJU3bWNi8S9Ws5tb8/H1GyeTUw8VUlsVyHPNMMHcHqTatbA3cotPqKgNi090mt4HfCLRt0G45VEmWKTiQf6hE7EugX6Iova/PMghQFQGhHyaD/jvT1745grtM6YytBnj5d3HSTEUoPb5V5pXRwxDwe9audBdg/QzCNYX69KJdXCER7ge7JULYjLdCdhUR7HUaFpHWHnapKcQm8MluueeHlRlftKO5lTCsp+HtmlqV2B0Lv6c3fz9LTvzEmEX4HygTBZqKl3Q/33bUg4FeD2cAWGwYMQMwwqDuAzBhe+L1CQyoqqDpk3LfW8RAoZJTaUc38+ZpggKpkX5GxPrbKW4CxVS6XoabfmT1WXG3HsWVyrrm55tn8mxkRyIUXccWwfDFplxCkzpIo1fv3l1cS454g7QjHF3pHMI62Q1Nh7FQ6/yIMXSGePWmDP788UACgXYXJB7RCtAddMllsFV7cjydHNi4p0rtBVoc++0pXRH9yVCjmFP+JO+yW2R9xMMXv+Fm3Yk+YJqqdJZAtX2BSvQnHEEg/HXqG1RG6DgbKpcdeD7dipmtfcucF27IpVEfpxl5wtQj2hzNsh4W4+rMJmeWSC9QlFpFRzXWiuMwhsgvi45HaMV2thy+i7e9e07sbpIR0fR4jjL1ut1Wtg6daHIeruYQeHN2UU6SkuqTEeGoYo3y1kXbQ82rqEoMKTaZa1KxnXS1LLU+dsuEZlIHocO/ii9SEdnQaUXI/brXaQK7AHSaySxv7GPct/uF+C7Xgh9I5ngM29At/dgW5VtP+L3cjficn9DJnK8fwPME8nTzKrb7QIi3gXTNCz+XGPgHZsncgVBw4KLwa8NHfl3LsdLMBG/ktuP//TM+JM4epQM4t5Npd1wncHU/E8m8hE3R4+WZt4kuxFnON35RCn0dGB5wivNIQ1cX36QiYR+PfcLwc6S3Q/2Pojq5UJ1EPizSb5gst1269Y0z/rd0Rctnre40+YizZum+R/KhI2R","sidebar_class_name":"get api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Delete a Message","permalink":"/docs/1backend/delete-message"},"next":{"title":"Add Thread","permalink":"/docs/1backend/add-thread"}}');var i=s(74848),d=s(28453),n=s(53746),c=s.n(n),o=s(56518),r=s.n(o),p=s(99972),l=s.n(p),g=s(25342),m=s.n(g),h=(s(44215),s(82223),s(24861));const u={id:"get-message",title:"Get Message",description:"Fetch information about a specific chat message by its ID",sidebar_label:"Get Message",hide_title:!0,hide_table_of_contents:!0,api:"eJylVltv2zYU/isEX7YBiuQ0K1D4ZXO3JPPWLsHcoA+JMRxTxxIbimTJIzueof8+HEmO7VgtUPTFl8Nz+c7tI7cyx6iC9qSdlWN5haRKoe3ShQpYJmDhahIgokell1oJVQKJCmOEAsViIzRFMf1dJtJ5DK3NNJdjWSC975RkIj0EqJAwRDm+376I2at1TjRLPFApE2mhQjmWfaxpLhMZ8HOtA+ZyTKHGREZVYgVyvJW08awcKWhbyKaZs3L0zkaMfP5qNOKv4dA5EmgTRayVwhiXtTEbEZCCxhVyXOUsoSX2AN4brdpEs0+R3WwPYPjAZSDdBcUnHSkewFs4ZxCsbJJdWqdGKiAQ5hMaSCuRS21wmsfTXK66A5HjUluMgkoUrCyACFRZoaVOuOtdCTHlihNWcTBUL4AQYMP/dd7lBJU3bWNi8S9Ws5tb8/H1GyeTUw8VUlsVyHPNMMHcHqTatbA3cotPqKgNi090mt4HfCLRt0G45VEmWKTiQf6hE7EugX6Iova/PMghQFQGhHyaD/jvT1745grtM6YytBnj5d3HSTEUoPb5V5pXRwxDwe9audBdg/QzCNYX69KJdXCER7ge7JULYjLdCdhUR7HUaFpHWHnapKcQm8MluueeHlRlftKO5lTCsp+HtmlqV2B0Lv6c3fz9LTvzEmEX4HygTBZqKl3Q/33bUg4FeD2cAWGwYMQMwwqDuAzBhe+L1CQyoqqDpk3LfW8RAoZJTaUc38+ZpggKpkX5GxPrbKW4CxVS6XoabfmT1WXG3HsWVyrrm55tn8mxkRyIUXccWwfDFplxCkzpIo1fv3l1cS454g7QjHF3pHMI62Q1Nh7FQ6/yIMXSGePWmDP788UACgXYXJB7RCtAddMllsFV7cjydHNi4p0rtBVoc++0pXRH9yVCjmFP+JO+yW2R9xMMXv+Fm3Yk+YJqqdJZAtX2BSvQnHEEg/HXqG1RG6DgbKpcdeD7dipmtfcucF27IpVEfpxl5wtQj2hzNsh4W4+rMJmeWSC9QlFpFRzXWiuMwhsgvi45HaMV2thy+i7e9e07sbpIR0fR4jjL1ut1Wtg6daHIeruYQeHN2UU6SkuqTEeGoYo3y1kXbQ82rqEoMKTaZa1KxnXS1LLU+dsuEZlIHocO/ii9SEdnQaUXI/brXaQK7AHSaySxv7GPct/uF+C7Xgh9I5ngM29At/dgW5VtP+L3cjficn9DJnK8fwPME8nTzKrb7QIi3gXTNCz+XGPgHZsncgVBw4KLwa8NHfl3LsdLMBG/ktuP//TM+JM4epQM4t5Npd1wncHU/E8m8hE3R4+WZt4kuxFnON35RCn0dGB5wivNIQ1cX36QiYR+PfcLwc6S3Q/2Pojq5UJ1EPizSb5gst1269Y0z/rd0Rctnre40+YizZum+R/KhI2R",sidebar_class_name:"get api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},b=void 0,f={},y=[];function H(e){const t={p:"p",...(0,d.R)(),...e.components};return(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(h.default,{as:"h1",className:"openapi__heading",children:"Get Message"}),"\n",(0,i.jsx)(c(),{method:"get",path:"/chat-svc/message/{messageId}",context:"endpoint"}),"\n",(0,i.jsx)(t.p,{children:"Fetch information about a specific chat message by its ID"}),"\n",(0,i.jsx)(h.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,i.jsx)(r(),{parameters:[{description:"Message ID",in:"path",name:"messageId",required:!0,schema:{type:"string"}}]}),"\n",(0,i.jsx)(l(),{title:"Body",body:void 0}),"\n",(0,i.jsx)(m(),{id:void 0,label:void 0,responses:{200:{description:"Message details successfully retrieved",content:{"application/json":{schema:{properties:{exists:{type:"boolean"},message:{properties:{createdAt:{type:"string"},fileIds:{description:"FileIds defines the file attachments the message has.",items:{type:"string"},type:"array"},id:{example:"msg_emSOPlW58o",type:"string"},meta:{additionalProperties:!0,type:"object"},text:{description:'Text content of the message eg. "Hi, what\'s up?"',type:"string"},threadId:{description:"ThreadId of the message.",example:"thr_emSOeEUWAg",type:"string"},updatedAt:{type:"string"},userId:{description:"UserId is the id of the user who wrote the message.\nFor AI messages this field is empty.",type:"string"}},required:["id","threadId"],type:"object"}},type:"object"}}}},400:{description:"Invalid JSON",content:{"application/json":{schema:{type:"string"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{type:"string"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{type:"string"}}}}}})]})}function M(e={}){const{wrapper:t}={...(0,d.R)(),...e.components};return t?(0,i.jsx)(t,{...e,children:(0,i.jsx)(H,{...e})}):H(e)}}}]);