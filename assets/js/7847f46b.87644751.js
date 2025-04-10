"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[2595],{33767:(e,s,t)=>{t.r(s),t.d(s,{assets:()=>P,contentTitle:()=>u,default:()=>b,frontMatter:()=>y,metadata:()=>a,toc:()=>N});const a=JSON.parse('{"id":"1backend/add-message","title":"Add Message","description":"Add a new message to a specific thread.","source":"@site/docs/1backend/add-message.api.mdx","sourceDirName":"1backend","slug":"/1backend/add-message","permalink":"/docs/1backend/add-message","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"add-message","title":"Add Message","description":"Add a new message to a specific thread.","sidebar_label":"Add Message","hide_title":true,"hide_table_of_contents":true,"api":"eJylVltv2zYU/isEX7YBipQsKzboZXO7dvPWLUGdoA+JMZyQxxJbiVRJyq5n6L8Ph5RsJdKKBXsxZPJcvnP7Dg9cohNWNV4ZzXO+kJIB07hjNToHBTJvGDDXoFAbJZgvLYJMecJNgxZIayl5zkHKP6ICT3gDFmr0aB3P7w5PPNwEC2z5M0+4ooMGfMkTrqFGnvPoYCl5wi1+apVFyXNvW0y4EyXWwPMD9/uGZJ23She869ZRGJ1/aeSeJITRHrWnT2iaSokANfvgCMNhZKqxFIhX6OhfH/T0QlgEj3LhZ9wnfKMqXMog+DjYN/GCSdwojY75EhkJM/AeRFmj9vFwyHYJjpKrPNZu1lV/ANbCnv4rSWL4GeqmoovaFX9hvbq6rt6/+MHwZGqhRh8iBykVwYTqehRqTHWvZB4+oPDBLX720/Bu8LNnfaqZ2TyKBIuU3fNfVcJ2JfivHGubH+/5HKBjyaf2+5sntilDp4h9aUPE+Pr2/aKYc9A28gvFax3aOee34ZypWCB1BEHybFcatrPG4yNc9/qNsWyxHA5IVTm2UVgFQ1g3fp9OIXbjZr+jmo6ysp6Uo5uedMnMIPcTyd7F2ZiMVHDrGqNdbPJvz8+naRiMuFYIdG7TVtWegZRIIP/zmD2j2UIw381BWeotVEqy31ZXfz7H+dNsRwcXMyXX0PrSWPX386Kbc/BiPgKPVkPFVmi3aNlra439f566hDsUrVV+H9j2JYJFu2h9yfO7NVGjh4KImL8qwbPVVlBH1ehLQ8TdmNAYgYRznokS/Jnbiiy2X3YY2rDL6iO/uwA+kntrK9LLKiOgKo3z+cXF5eX3nBwPuFYEP7bYGN1k2vcNsvte5J6zjakqs0PJHvZhBYFABloybz6iZiBiK7ONNXWYQhpYio+9NYXSDLVsjNI+HRZNiSDRnlbNoq91yPVpKKFRv+M+TBnl/91psbweOOfRohithqEuo41wN5ytI1dPCbrn4yPJnoyciHFKciNOOykMVDYmFqU3ZtiIIEKLYQ2KquagQveTU7poK/DW6FSYepSf6yVbtU1jLHVILHTpfZNn2cUDiI+oJSlkfEo+yzMNXm2R1UpYQ/2iBDrWVOA3xtZUkkoJ1C5kcPD3y/Vbtr1Mzx95c3mW7Xa7tNBtamyR9Xoug6Kpzi7T87T0dRV3lK3d1WYVvZ3Auh0UBdpUmSyIZJRa5cPyuHgZA+EJp5aO8M/Ty/T8zIr08gXZpRGpQY+RnsiVP4l99PR4xmOqbz1qgKypQOmwmCgHh3407/gwmsfG4AnPR8+loSXXCac5JJXD4QEc3tqq6+j4U4uWSGKd8C1YBQ+UAnqgKUffkucbqBx+IaKv3/UL5Bs2fsfNwh/GSe8puVC19I8n/CPux++8bt0lw2gSmHj9Kro8I04YqU/IsUsGjYUQ2Pgvyo557/pqdcMT/tC/F2sjScfCjrYk7CJUE1IQiCucHXgFumjD4PNok0YMejo7EQhBSoYPimo2GU8JKAZCvxTWrMrhEOmp647y8epfNY6sF6WpNuuu6/4Bxr0Xkw==","sidebar_class_name":"post api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Update Thread","permalink":"/docs/1backend/update-thread"},"next":{"title":"List Messages","permalink":"/docs/1backend/get-messages"}}');var i=t(74848),n=t(28453),d=t(53746),o=t.n(d),r=t(56518),c=t.n(r),p=t(99972),h=t.n(p),l=t(25342),g=t.n(l),m=(t(44215),t(82223),t(24861));const y={id:"add-message",title:"Add Message",description:"Add a new message to a specific thread.",sidebar_label:"Add Message",hide_title:!0,hide_table_of_contents:!0,api:"eJylVltv2zYU/isEX7YBipQsKzboZXO7dvPWLUGdoA+JMZyQxxJbiVRJyq5n6L8Ph5RsJdKKBXsxZPJcvnP7Dg9cohNWNV4ZzXO+kJIB07hjNToHBTJvGDDXoFAbJZgvLYJMecJNgxZIayl5zkHKP6ICT3gDFmr0aB3P7w5PPNwEC2z5M0+4ooMGfMkTrqFGnvPoYCl5wi1+apVFyXNvW0y4EyXWwPMD9/uGZJ23She869ZRGJ1/aeSeJITRHrWnT2iaSokANfvgCMNhZKqxFIhX6OhfH/T0QlgEj3LhZ9wnfKMqXMog+DjYN/GCSdwojY75EhkJM/AeRFmj9vFwyHYJjpKrPNZu1lV/ANbCnv4rSWL4GeqmoovaFX9hvbq6rt6/+MHwZGqhRh8iBykVwYTqehRqTHWvZB4+oPDBLX720/Bu8LNnfaqZ2TyKBIuU3fNfVcJ2JfivHGubH+/5HKBjyaf2+5sntilDp4h9aUPE+Pr2/aKYc9A28gvFax3aOee34ZypWCB1BEHybFcatrPG4yNc9/qNsWyxHA5IVTm2UVgFQ1g3fp9OIXbjZr+jmo6ysp6Uo5uedMnMIPcTyd7F2ZiMVHDrGqNdbPJvz8+naRiMuFYIdG7TVtWegZRIIP/zmD2j2UIw381BWeotVEqy31ZXfz7H+dNsRwcXMyXX0PrSWPX386Kbc/BiPgKPVkPFVmi3aNlra439f566hDsUrVV+H9j2JYJFu2h9yfO7NVGjh4KImL8qwbPVVlBH1ehLQ8TdmNAYgYRznokS/Jnbiiy2X3YY2rDL6iO/uwA+kntrK9LLKiOgKo3z+cXF5eX3nBwPuFYEP7bYGN1k2vcNsvte5J6zjakqs0PJHvZhBYFABloybz6iZiBiK7ONNXWYQhpYio+9NYXSDLVsjNI+HRZNiSDRnlbNoq91yPVpKKFRv+M+TBnl/91psbweOOfRohithqEuo41wN5ytI1dPCbrn4yPJnoyciHFKciNOOykMVDYmFqU3ZtiIIEKLYQ2KquagQveTU7poK/DW6FSYepSf6yVbtU1jLHVILHTpfZNn2cUDiI+oJSlkfEo+yzMNXm2R1UpYQ/2iBDrWVOA3xtZUkkoJ1C5kcPD3y/Vbtr1Mzx95c3mW7Xa7tNBtamyR9Xoug6Kpzi7T87T0dRV3lK3d1WYVvZ3Auh0UBdpUmSyIZJRa5cPyuHgZA+EJp5aO8M/Ty/T8zIr08gXZpRGpQY+RnsiVP4l99PR4xmOqbz1qgKypQOmwmCgHh3407/gwmsfG4AnPR8+loSXXCac5JJXD4QEc3tqq6+j4U4uWSGKd8C1YBQ+UAnqgKUffkucbqBx+IaKv3/UL5Bs2fsfNwh/GSe8puVC19I8n/CPux++8bt0lw2gSmHj9Kro8I04YqU/IsUsGjYUQ2Pgvyo557/pqdcMT/tC/F2sjScfCjrYk7CJUE1IQiCucHXgFumjD4PNok0YMejo7EQhBSoYPimo2GU8JKAZCvxTWrMrhEOmp647y8epfNY6sF6WpNuuu6/4Bxr0Xkw==",sidebar_class_name:"post api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},u=void 0,P={},N=[];function f(e){const s={p:"p",...(0,n.R)(),...e.components};return(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(m.default,{as:"h1",className:"openapi__heading",children:"Add Message"}),"\n",(0,i.jsx)(o(),{method:"post",path:"/chat-svc/thread/{threadId}/message",context:"endpoint"}),"\n",(0,i.jsx)(s.p,{children:"Add a new message to a specific thread."}),"\n",(0,i.jsx)(m.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,i.jsx)(c(),{parameters:[{description:"Thread ID",in:"path",name:"threadId",required:!0,schema:{type:"string"}}]}),"\n",(0,i.jsx)(h(),{title:"Body",body:{content:{"application/json":{schema:{properties:{message:{properties:{createdAt:{type:"string"},fileIds:{description:"FileIds defines the file attachments the message has.",items:{type:"string"},type:"array"},id:{example:"msg_emSOPlW58o",type:"string"},meta:{additionalProperties:!0,type:"object"},text:{description:'Text content of the message eg. "Hi, what\'s up?"',type:"string"},threadId:{description:"ThreadId of the message.",example:"thr_emSOeEUWAg",type:"string"},updatedAt:{type:"string"},userId:{description:"UserId is the id of the user who wrote the message.\nFor AI messages this field is empty.",type:"string"}},required:["id","threadId"],type:"object"}},type:"object"}}},description:"Add Message Request",required:!0}}),"\n",(0,i.jsx)(g(),{id:void 0,label:void 0,responses:{200:{description:"Message successfully added",content:{"application/json":{schema:{additionalProperties:!0,type:"object"}}}},400:{description:"Invalid JSON",content:{"application/json":{schema:{type:"string"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{type:"string"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{type:"string"}}}}}})]})}function b(e={}){const{wrapper:s}={...(0,n.R)(),...e.components};return s?(0,i.jsx)(s,{...e,children:(0,i.jsx)(f,{...e})}):f(e)}}}]);