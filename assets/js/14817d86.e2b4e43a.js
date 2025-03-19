"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[3318],{21211:(e,s,i)=>{i.r(s),i.d(s,{assets:()=>k,contentTitle:()=>v,default:()=>x,frontMatter:()=>h,metadata:()=>t,toc:()=>b});const t=JSON.parse('{"id":"1backend/save-self","title":"Save User Profile","description":"Save user\'s own profile information.","source":"@site/docs/1backend/save-self.api.mdx","sourceDirName":"1backend","slug":"/1backend/save-self","permalink":"/docs/1backend/save-self","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"save-self","title":"Save User Profile","description":"Save user\'s own profile information.","sidebar_label":"Save User Profile","hide_title":true,"hide_table_of_contents":true,"api":"eJy9VcFy2zYQ/RXMXprOUKRsNzMtT5HbTEdNpvZY8cnRYUUuKcQggACgFFXDf+8sIVq0pWbcSy4SASywbx/2PeyhJF84aYM0GnJY4IZE68n95IXZamGdqaQiIXVlXIMclUICxpLrB/MScvC4oQWpChKw6LChQM5D/rB/cfi9Jyfmf0ACkocWwxoS0NgQ5MBJ5yUk4OhrKx2VkAfXUgK+WFODkO8h7CxH+uCkrqHrljGYfLg25Y4jCqMD6cCfaK2SRQ8y++I5/350lHVcQpDkeRQhnCRIwKu2Ppc5GWbM6gsVATqeOsPk7YG+uwjzpLqOJ7w12kcgl9Mp/z0/6eYDJK+v7ARYl8Av546d6w0qWYq/Fjd//58Ez6kj54x7JUU9kotTJPca27A2Tv5D5Q9D8vY8J4GcRiUW5DbkxPv+zB8DiduNitbJsOu1c03oyM3asIb8YcnNHrBmWUUdLTYFLBNoKKwNi9C2odcfh0PGapr4TZH5qEvf1xNF2TrFIZkyBaq18SF/++vl1QVwigHBgiuKRYxxvOTr086S+HwI+QyiMkqZLZVitRMovMWCBOpSBPNIWmARe19UzjQirEkMlYiPppZakC6tkTqkg0WsCUtyR5OYHfqkpx+eWEQrP9Cu55Wv5O5oCu+/YWMVHUU+3MGg7ZHY2eMGE8Giv2xqUKre4RT5d17qulUYnNFpYZoRrNu5WLTWGsd3EPldh2DzLLtYYfFIuuQNGZy4xEy0WlaSGYtxoqQNKWMb0kFYhYF9V1TGiUYWzvA9yoL8ZIWeSjGbi1E/euZNyYK0H1c8+/P2o9hcpdNn2HyeZdvtNq11mxpXZ4d9PsPaqslVOk3XoVGMOJBr/E21iKmPpfkt1jW5VJqsD8n4QmRgtuHiOpYDCXDfxWKn6VU6nbgivfyNz7XGhwb1CGnvmX1PHIwTXvA1cvjXPlWHFgn0LWRWodScuqdhf1DLAwxq6YWiKtYV64KX9ntm+t6pruPpry05lucygQ06iSuulh866fm7hLxC5ek7sN/cHV6An8XxPTwLcmhuvWMWUbU8ggQeaXd8L7tllwwyYSBx8feYbsL6HG0+8a4uGXbMioJs+G7s2G1u7z9BAqvDs9uYkrc43PILh9uI0/TV9x7Sz+1Boa5brDk2HsmKxYOzHLXMiJLhg4s6y8RLL4h18C9XdXbLfh+douue4uPSf+54MqAYzRez7LruX9pNJ94=","sidebar_class_name":"put api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Assign Permissions","permalink":"/docs/1backend/assign-permissions"},"next":{"title":"Create a New User","permalink":"/docs/1backend/create-user"}}');var r=i(74848),n=i(28453),a=i(53746),o=i.n(a),c=i(56518),p=i.n(c),d=i(99972),l=i.n(d),u=i(25342),f=i.n(u),Y=(i(44215),i(82223),i(24861));const h={id:"save-self",title:"Save User Profile",description:"Save user's own profile information.",sidebar_label:"Save User Profile",hide_title:!0,hide_table_of_contents:!0,api:"eJy9VcFy2zYQ/RXMXprOUKRsNzMtT5HbTEdNpvZY8cnRYUUuKcQggACgFFXDf+8sIVq0pWbcSy4SASywbx/2PeyhJF84aYM0GnJY4IZE68n95IXZamGdqaQiIXVlXIMclUICxpLrB/MScvC4oQWpChKw6LChQM5D/rB/cfi9Jyfmf0ACkocWwxoS0NgQ5MBJ5yUk4OhrKx2VkAfXUgK+WFODkO8h7CxH+uCkrqHrljGYfLg25Y4jCqMD6cCfaK2SRQ8y++I5/350lHVcQpDkeRQhnCRIwKu2Ppc5GWbM6gsVATqeOsPk7YG+uwjzpLqOJ7w12kcgl9Mp/z0/6eYDJK+v7ARYl8Av546d6w0qWYq/Fjd//58Ez6kj54x7JUU9kotTJPca27A2Tv5D5Q9D8vY8J4GcRiUW5DbkxPv+zB8DiduNitbJsOu1c03oyM3asIb8YcnNHrBmWUUdLTYFLBNoKKwNi9C2odcfh0PGapr4TZH5qEvf1xNF2TrFIZkyBaq18SF/++vl1QVwigHBgiuKRYxxvOTr086S+HwI+QyiMkqZLZVitRMovMWCBOpSBPNIWmARe19UzjQirEkMlYiPppZakC6tkTqkg0WsCUtyR5OYHfqkpx+eWEQrP9Cu55Wv5O5oCu+/YWMVHUU+3MGg7ZHY2eMGE8Giv2xqUKre4RT5d17qulUYnNFpYZoRrNu5WLTWGsd3EPldh2DzLLtYYfFIuuQNGZy4xEy0WlaSGYtxoqQNKWMb0kFYhYF9V1TGiUYWzvA9yoL8ZIWeSjGbi1E/euZNyYK0H1c8+/P2o9hcpdNn2HyeZdvtNq11mxpXZ4d9PsPaqslVOk3XoVGMOJBr/E21iKmPpfkt1jW5VJqsD8n4QmRgtuHiOpYDCXDfxWKn6VU6nbgivfyNz7XGhwb1CGnvmX1PHIwTXvA1cvjXPlWHFgn0LWRWodScuqdhf1DLAwxq6YWiKtYV64KX9ntm+t6pruPpry05lucygQ06iSuulh866fm7hLxC5ek7sN/cHV6An8XxPTwLcmhuvWMWUbU8ggQeaXd8L7tllwwyYSBx8feYbsL6HG0+8a4uGXbMioJs+G7s2G1u7z9BAqvDs9uYkrc43PILh9uI0/TV9x7Sz+1Boa5brDk2HsmKxYOzHLXMiJLhg4s6y8RLL4h18C9XdXbLfh+douue4uPSf+54MqAYzRez7LruX9pNJ94=",sidebar_class_name:"put api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},v=void 0,k={},b=[];function m(e){const s={p:"p",...(0,n.R)(),...e.components};return(0,r.jsxs)(r.Fragment,{children:[(0,r.jsx)(Y.default,{as:"h1",className:"openapi__heading",children:"Save User Profile"}),"\n",(0,r.jsx)(o(),{method:"put",path:"/user-svc/self",context:"endpoint"}),"\n",(0,r.jsx)(s.p,{children:"Save user's own profile information."}),"\n",(0,r.jsx)(Y.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,r.jsx)(p(),{parameters:[{description:"User ID",in:"path",name:"userId",required:!0,schema:{type:"string"}}]}),"\n",(0,r.jsx)(l(),{title:"Body",body:{content:{"application/json":{schema:{properties:{name:{type:"string"},slug:{type:"string"}},type:"object"}}},description:"Save Profile Request",required:!0}}),"\n",(0,r.jsx)(f(),{id:void 0,label:void 0,responses:{200:{description:"OK",content:{"application/json":{schema:{type:"object"}}}},400:{description:"Invalid JSON",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}}}})]})}function x(e={}){const{wrapper:s}={...(0,n.R)(),...e.components};return s?(0,r.jsx)(s,{...e,children:(0,r.jsx)(m,{...e})}):m(e)}}}]);