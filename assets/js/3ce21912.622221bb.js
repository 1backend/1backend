"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[5815],{76671:(e,t,a)=>{a.r(t),a.d(t,{assets:()=>f,contentTitle:()=>h,default:()=>y,frontMatter:()=>m,metadata:()=>n,toc:()=>w});const n=JSON.parse('{"id":"1backend/get-download","title":"Get a Download","description":"Get a download by URL.","source":"@site/docs/1backend/get-download.api.mdx","sourceDirName":"1backend","slug":"/1backend/get-download","permalink":"/docs/1backend/get-download","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"get-download","title":"Get a Download","description":"Get a download by URL.","sidebar_label":"Get a Download","hide_title":true,"hide_table_of_contents":true,"api":"eJytVVFv2zYQ/isEnzZAkZykHTo9zUG7IGiwBnXzlBjoWTrLbCiSJU92HUP/fTjKsuXYKTZsL7ak+3j33d13x40sMRReOVLWyFxeIwkQpV0ZbaEUs7W4/3ybPppH8xm/N8pjELRA8XWuNJ6FZZH30HypcPVVOPS1CkFZk8pEWoce2PNNKXNZIb3fomUiHXiokdAHmT9sXtBovJaJVPzogBYykQZq3Bl8x6WUOfkGExmKBdYg842ktWNYIK9MJdt2yuDgrAkY2H4xGvHfYbRPH2UiC2sIDbEVnNOqiMSzb4Ehm0EI5zktUp3DPv9jS+ERCMsxneCV7M5hebWmrasDUu8PAQJ/qEBBkBVhYVexC7s2OW8rjyEIZQShr4Ow84gwTT1Dz2+z6AS0RyjXYh+d2zS3vgaSuVSGfnsjk56tMoQVeqaL3lt/MhFWwl+xOa8Y77iBrxkn6hmPc/9zaxGqk9u80VrMlQE9YC74vAjqGf9hDqo8SaMv3sDYlY2NgYCacPJc48qf9JeFekKPO1529g0LiqWNjR2AZ9ZqBBPRe6U/9MDpkY+WkW9G58eFvDfQ0MJ69Yzlv9H4S9oc4O2p0bkxhJ77MkG/RC8+RJ38p0hcdSwar2gdN8MVgkc/blhFD1OeaIKKl0aUiZgsCy5JjbSw2y0TtwvDZdbvqazXTbZpvG4lx2DC3fKJzZJZpm0BemED5W/fXVyeSw7Wc5kw5W5Qh4xeFuTL2qF43EIepZhbre0K4yoFERwUKMCUguwTGgFF114x97aOUr8P6DkncWsrZQSa0lllKO3X4QKhRL9fiONtf2N997IHpz7iOkpImbmN+8gagiK2BGtQnHEAjeGPoEzVaCBvTVrYeuD77kZMGues55J2RVoQuTzLzmdQPKEp+UAW99lBFcY3ZwZILVHUqvCWa60KDMJpIJ5UTkerAk2I09/Hu767FcvLdHQQLeRZtlqt0so0qfVVtj0XMqicPrtMR+mCas0c4ur7NJ900fZkwwqqCn2qbBYhGddJkWbI+VWXiEwky6GjP0ov09GZL9KL3+OGsIFqMAOm3TU5uM0O0t/s5f8/X6jb7hL+oMxpUGawajrJP8jek9xfMjKROaOmiWR1M2qzmUHAe6/blj9/b9DzuE0TuQSvYMbF4WtZBX4uZT4HHfAnif6yzaf8VXSX9EmuvTzNmgsOuuE3mcgnXG9v93baJr3KmUFnGRcFOhqcOdoq7XAJXH/4IhMJzcHNA06xs6R/YO8n+bycqY4C/7bJK0c2m27i2naH70yvntgNcofm8kzbtv0bLuxIfA==","sidebar_class_name":"get api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Download a File","permalink":"/docs/1backend/download-file"},"next":{"title":"Pause a Download","permalink":"/docs/1backend/pause-download"}}');var o=a(74848),s=a(28453),i=a(53746),d=a.n(i),r=a(56518),l=a.n(r),c=a(99972),p=a.n(c),b=a(25342),u=a.n(b),k=(a(44215),a(82223),a(24861));const m={id:"get-download",title:"Get a Download",description:"Get a download by URL.",sidebar_label:"Get a Download",hide_title:!0,hide_table_of_contents:!0,api:"eJytVVFv2zYQ/isEnzZAkZykHTo9zUG7IGiwBnXzlBjoWTrLbCiSJU92HUP/fTjKsuXYKTZsL7ak+3j33d13x40sMRReOVLWyFxeIwkQpV0ZbaEUs7W4/3ybPppH8xm/N8pjELRA8XWuNJ6FZZH30HypcPVVOPS1CkFZk8pEWoce2PNNKXNZIb3fomUiHXiokdAHmT9sXtBovJaJVPzogBYykQZq3Bl8x6WUOfkGExmKBdYg842ktWNYIK9MJdt2yuDgrAkY2H4xGvHfYbRPH2UiC2sIDbEVnNOqiMSzb4Ehm0EI5zktUp3DPv9jS+ERCMsxneCV7M5hebWmrasDUu8PAQJ/qEBBkBVhYVexC7s2OW8rjyEIZQShr4Ow84gwTT1Dz2+z6AS0RyjXYh+d2zS3vgaSuVSGfnsjk56tMoQVeqaL3lt/MhFWwl+xOa8Y77iBrxkn6hmPc/9zaxGqk9u80VrMlQE9YC74vAjqGf9hDqo8SaMv3sDYlY2NgYCacPJc48qf9JeFekKPO1529g0LiqWNjR2AZ9ZqBBPRe6U/9MDpkY+WkW9G58eFvDfQ0MJ69Yzlv9H4S9oc4O2p0bkxhJ77MkG/RC8+RJ38p0hcdSwar2gdN8MVgkc/blhFD1OeaIKKl0aUiZgsCy5JjbSw2y0TtwvDZdbvqazXTbZpvG4lx2DC3fKJzZJZpm0BemED5W/fXVyeSw7Wc5kw5W5Qh4xeFuTL2qF43EIepZhbre0K4yoFERwUKMCUguwTGgFF114x97aOUr8P6DkncWsrZQSa0lllKO3X4QKhRL9fiONtf2N997IHpz7iOkpImbmN+8gagiK2BGtQnHEAjeGPoEzVaCBvTVrYeuD77kZMGues55J2RVoQuTzLzmdQPKEp+UAW99lBFcY3ZwZILVHUqvCWa60KDMJpIJ5UTkerAk2I09/Hu767FcvLdHQQLeRZtlqt0so0qfVVtj0XMqicPrtMR+mCas0c4ur7NJ900fZkwwqqCn2qbBYhGddJkWbI+VWXiEwky6GjP0ov09GZL9KL3+OGsIFqMAOm3TU5uM0O0t/s5f8/X6jb7hL+oMxpUGawajrJP8jek9xfMjKROaOmiWR1M2qzmUHAe6/blj9/b9DzuE0TuQSvYMbF4WtZBX4uZT4HHfAnif6yzaf8VXSX9EmuvTzNmgsOuuE3mcgnXG9v93baJr3KmUFnGRcFOhqcOdoq7XAJXH/4IhMJzcHNA06xs6R/YO8n+bycqY4C/7bJK0c2m27i2naH70yvntgNcofm8kzbtv0bLuxIfA==",sidebar_class_name:"get api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},h=void 0,f={},w=[];function g(e){const t={code:"code",p:"p",...(0,s.R)(),...e.components};return(0,o.jsxs)(o.Fragment,{children:[(0,o.jsx)(k.default,{as:"h1",className:"openapi__heading",children:"Get a Download"}),"\n",(0,o.jsx)(d(),{method:"get",path:"/file-svc/download/{url}",context:"endpoint"}),"\n",(0,o.jsx)(t.p,{children:"Get a download by URL."}),"\n",(0,o.jsxs)(t.p,{children:["Requires the ",(0,o.jsx)(t.code,{children:"file-svc:download:view"})," permission."]}),"\n",(0,o.jsx)(k.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,o.jsx)(l(),{parameters:[{description:"url",in:"path",name:"url",required:!0,schema:{type:"string"}}]}),"\n",(0,o.jsx)(p(),{title:"Body",body:void 0}),"\n",(0,o.jsx)(u(),{id:void 0,label:void 0,responses:{200:{description:"OK",content:{"application/json":{schema:{properties:{download:{properties:{createdAt:{type:"string"},downloadedBytes:{description:"DownloadedBytes exists to show the download progress in terms of the number of bytes already downloaded.",format:"int64",type:"integer"},error:{type:"string"},fileName:{type:"string"},filePath:{type:"string"},fileSize:{description:"FileSize is the full final downloaded file size.",format:"int64",type:"integer"},id:{type:"string"},progress:{type:"number"},status:{type:"string"},updatedAt:{type:"string"},url:{type:"string"}},type:"object"},exists:{type:"boolean"}},required:["exists"],type:"object"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{type:"string"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{type:"string"}}}}}})]})}function y(e={}){const{wrapper:t}={...(0,s.R)(),...e.components};return t?(0,o.jsx)(t,{...e,children:(0,o.jsx)(g,{...e})}):g(e)}}}]);