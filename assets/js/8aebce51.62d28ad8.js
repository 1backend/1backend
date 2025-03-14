"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[7385],{80677:(e,t,a)=>{a.r(t),a.d(t,{assets:()=>m,contentTitle:()=>b,default:()=>M,frontMatter:()=>k,metadata:()=>s,toc:()=>y});const s=JSON.parse('{"id":"1backend/make-default","title":"Make a Model Default","description":"Sets a model as the default model \u2014 when prompts are sent without a Model ID, the default model is used.","source":"@site/docs/1backend/make-default.api.mdx","sourceDirName":"1backend","slug":"/1backend/make-default","permalink":"/docs/1backend/make-default","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"make-default","title":"Make a Model Default","description":"Sets a model as the default model \u2014 when prompts are sent without a Model ID, the default model is used.","sidebar_label":"Make a Model Default","hide_title":true,"hide_table_of_contents":true,"api":"eJy9VU1v3DYQ/SvEnFqAltZ2A7Q61UGKYpukNrrxydnDLDW7YkyRDEntdisIyI/IL8wvKUbSftheB8jFl11JnBm+eXxv2EJJUQXtk3YWCphRigJF7UoyAqNIFYmSltiYNH789uWr2FRkhQ+u9hwdSESySWx0qlyTBIr3feT0jTyRr6NoIpUZSHCeAvLG0xIKqPGe3gyhIMFjwJoShQjFXfsI5a4+SND87jFVIMFiTVyIV6clSAj0udGBSihSaEhCVBXVCEULaes5NKag7Qq6bs7B0TsbKfL6xWTCfw+3vX4LEpSziWziVfTeaNV3kH+KHNI+3cItPpFK0HVdJ+GXU2Wndo1Gl+Kv2fXfP7KBD0xg0gNiCsGFU63JZ5CcP0Vya7FJlQv6PypfDMmr05wkChaNmFFYUxB/9DVfBlInIZJqgk7bXnuvCQOFqyZVUNzNWSoJVyzLUYeztYK5hJpS5VjIvhkEzPGQ92o8i2s1POXtKM8uZ8GflXvFx77TQe5NMJybG6fQVC6m4tWvF5fnwJvvsM2416G9Y4SPmfyw9SQ+jiEfQSydMW5DpVhsBYroUZFAW4rk7skKVINjxDK4unfvbaTALYp3bqWtIFt6p23Kdt6rCEsKB/ddjQrqDwb2/KLXb2nbM67t0jFOPktU/VlSjZo7jmgo/h61XTUGU3A2U64+qn0zFbPGexeYsIGkKiVf5Pn5AtU92ZITcujkIxauhHJ17awYw8TSBbF1TRBXU3GkpPjty9daq+D4NLSieLbASGXP0KLRJonkRFRoiAkwWpGNxB3sEP55806sL7PJA3yxyPPNZpOtbJO5sMrHvJjjypuzy2ySVak2jDpRqOP1cjbsfmgvbnC1opBpl/chOTOrk+GQ89dDTyCBBTQ0PMkus8lZUNnFb1zXu5hqtEdI3+M97Uf1Ye4+oK09+O1lb4ZRNIn+Tbk3qC330PPZjsa6g72xQA7PIKE4zP4H7ppLYBdxWtvyid4G03X8+XNDgW0+l7DGoHHBlPKFoyM/l1As0UT6DjE//TPeMj+Lo3vpZAc7L9gtnxWaht9Awj1tj+6tbt7Jna0YyrB6pRT5dJT3ZP51x0Po5vYDSMBxJBxMyMXk7oGrn8T02MQDBP7t5DMpbTtYvOv28cPSsxn7yTFEM0Xzruv+B+HX7/4=","sidebar_class_name":"put api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Get a Model","permalink":"/docs/1backend/get-model"},"next":{"title":"Start a Model","permalink":"/docs/1backend/start-model"}}');var n=a(74848),o=a(28453),d=a(53746),l=a.n(d),i=a(56518),r=a.n(i),p=a(99972),c=a.n(p),u=a(25342),h=a.n(u),f=(a(44215),a(82223),a(24861));const k={id:"make-default",title:"Make a Model Default",description:"Sets a model as the default model \u2014 when prompts are sent without a Model ID, the default model is used.",sidebar_label:"Make a Model Default",hide_title:!0,hide_table_of_contents:!0,api:"eJy9VU1v3DYQ/SvEnFqAltZ2A7Q61UGKYpukNrrxydnDLDW7YkyRDEntdisIyI/IL8wvKUbSftheB8jFl11JnBm+eXxv2EJJUQXtk3YWCphRigJF7UoyAqNIFYmSltiYNH789uWr2FRkhQ+u9hwdSESySWx0qlyTBIr3feT0jTyRr6NoIpUZSHCeAvLG0xIKqPGe3gyhIMFjwJoShQjFXfsI5a4+SND87jFVIMFiTVyIV6clSAj0udGBSihSaEhCVBXVCEULaes5NKag7Qq6bs7B0TsbKfL6xWTCfw+3vX4LEpSziWziVfTeaNV3kH+KHNI+3cItPpFK0HVdJ+GXU2Wndo1Gl+Kv2fXfP7KBD0xg0gNiCsGFU63JZ5CcP0Vya7FJlQv6PypfDMmr05wkChaNmFFYUxB/9DVfBlInIZJqgk7bXnuvCQOFqyZVUNzNWSoJVyzLUYeztYK5hJpS5VjIvhkEzPGQ92o8i2s1POXtKM8uZ8GflXvFx77TQe5NMJybG6fQVC6m4tWvF5fnwJvvsM2416G9Y4SPmfyw9SQ+jiEfQSydMW5DpVhsBYroUZFAW4rk7skKVINjxDK4unfvbaTALYp3bqWtIFt6p23Kdt6rCEsKB/ddjQrqDwb2/KLXb2nbM67t0jFOPktU/VlSjZo7jmgo/h61XTUGU3A2U64+qn0zFbPGexeYsIGkKiVf5Pn5AtU92ZITcujkIxauhHJ17awYw8TSBbF1TRBXU3GkpPjty9daq+D4NLSieLbASGXP0KLRJonkRFRoiAkwWpGNxB3sEP55806sL7PJA3yxyPPNZpOtbJO5sMrHvJjjypuzy2ySVak2jDpRqOP1cjbsfmgvbnC1opBpl/chOTOrk+GQ89dDTyCBBTQ0PMkus8lZUNnFb1zXu5hqtEdI3+M97Uf1Ye4+oK09+O1lb4ZRNIn+Tbk3qC330PPZjsa6g72xQA7PIKE4zP4H7ppLYBdxWtvyid4G03X8+XNDgW0+l7DGoHHBlPKFoyM/l1As0UT6DjE//TPeMj+Lo3vpZAc7L9gtnxWaht9Awj1tj+6tbt7Jna0YyrB6pRT5dJT3ZP51x0Po5vYDSMBxJBxMyMXk7oGrn8T02MQDBP7t5DMpbTtYvOv28cPSsxn7yTFEM0Xzruv+B+HX7/4=",sidebar_class_name:"put api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},b=void 0,m={},y=[];function j(e){const t={p:"p",...(0,o.R)(),...e.components};return(0,n.jsxs)(n.Fragment,{children:[(0,n.jsx)(f.default,{as:"h1",className:"openapi__heading",children:"Make a Model Default"}),"\n",(0,n.jsx)(l(),{method:"put",path:"/model-svc/model/{modelId}/make-default",context:"endpoint"}),"\n",(0,n.jsx)(t.p,{children:"Sets a model as the default model \u2014 when prompts are sent without a Model ID, the default model is used."}),"\n",(0,n.jsx)(f.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,n.jsx)(r(),{parameters:[{description:"Model ID",in:"path",name:"modelId",required:!0,schema:{type:"string"}}]}),"\n",(0,n.jsx)(c(),{title:"Body",body:void 0}),"\n",(0,n.jsx)(h(),{id:void 0,label:void 0,responses:{200:{description:"OK",content:{"application/json":{schema:{type:"object"}}}},400:{description:"Invalid JSON",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}}}})]})}function M(e={}){const{wrapper:t}={...(0,o.R)(),...e.components};return t?(0,n.jsx)(t,{...e,children:(0,n.jsx)(j,{...e})}):j(e)}}}]);