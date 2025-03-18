"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[7724],{90321:(e,t,i)=>{i.r(t),i.d(t,{assets:()=>M,contentTitle:()=>N,default:()=>y,frontMatter:()=>b,metadata:()=>a,toc:()=>g});const a=JSON.parse('{"id":"1backend/build-image","title":"Build an Image","description":"Builds a Docker image with the specified parameters.","source":"@site/docs/1backend/build-image.api.mdx","sourceDirName":"1backend","slug":"/1backend/build-image","permalink":"/docs/1backend/build-image","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"build-image","title":"Build an Image","description":"Builds a Docker image with the specified parameters.","sidebar_label":"Build an Image","hide_title":true,"hide_table_of_contents":true,"api":"eJy9Vt1u2zYUfhWC14rkJCuw6WpJWwxZiyaIm6vEQI+pY5kNRTIkZcczBOwh9oR7kuGQUmzF7rDd5MaQqfPznZ/vo7a8Qi+ctEEazUt+2UpVeQbsgxGP6JhsoEa2lmHJwhKZtyjkQmLFLDhoMKDz+YN+0Lf41EqHPlp9E0YHkBrdiV+JMsYo5xT5G7PoGum9NDrnGTcWHVDqq4qXPJpckTXPuMOnFn24NNWGl1tOIVEHegRrlRTRrfjuCfaWe7HEBujJOgoaJPoXt+dwA2FJf8fFvt+9ZDJBV0aAYpZOgoknERTr4/CM4zM0ViEvORUQNpYefXBS17zLeBX7tpAKj+f8MHr/47Q7u1HO0fFBcg0NHqb8Ag0OiciCmUV8TrMNJpU4SqNrqZ9LBQF9OEzUpelIhxUv70c97jHMXnzM/DuKwDtyOrJqLI6b3aZh8/3AwbUYM3lrtE/jPJtMDuu7/sSz/74fB7i6jP90LOyVXoGSFft9ev3l/yQYLyA6Z9xe3r0eHkdyeojkTkMblsbJP7B6MyTvjvckoNOg2BTdCh37GGO+DaQu4x5F62TY8PJ+yy8RHLqLllh2P+to5aD2tJDvB/lh05WgXWwwLA0pjG1px2xkJi9GMlXIXnh8LM3HHK1TZFhEfi6ND+W7n8/OTzllG8BMqbhUzz6k1637urHIHnqTB84WRimzxorNNwyYtyCQga5YMI+oGYhEA7ZwpolsvfOpHvbZ1FIz1JU1UgcSIUnxlwgVuoGAJb/oVyZOYsdhsPITbmKLaTq3O5X9OLD/lWpGnXsta2Mh6lOOVKMjXAszhAMRtwMbkNRSDwr9r17qulUQnNG5MM0e+JsrNm2tNY7mlaawDMGWRXE6B/GIuiKHgh+oygVrdbqhejtW4QqVsQ3qwKyCsDCuYQvj2Ma0jl1csb2N9X//+VcjhTO0BFKgP5mDxyoOhlQykFx6AQqpKUoK1D52bMD9281ntjrPJyPUviyK9Xqd17rNjauL3s8XUFt1cp5P8mVoFNUS0DX+ejFN2XdF+zXUNbpcmiKaFDRQGaJWn16mQnnGaW9TGyb5eT45cSI/+4XiWuNDA3oPaVJf0Gy4b0dt3Ltv3/qToN9T2r/CKpCa8MdebnviphvnJRgRINYwyzhRlAy2W5rbnVNdR8dPLToSjVnGV+AkzKlx97MuG1hDXH/EzfBNoMMJ0ZU6CqpNtHmlal02eFwIgTb8q+2+BN3cfeUZn/efNY2pyMXBmu4+WPOSx8+iuIxkEM+2XIGuWyqy5CkkERh6odlRmxBlwwMVNbzSmz2Ar6Uh1UG/VNVRl+02CUfXvdinVz/0eNGjZE1znHVd9w/vkJ7n","sidebar_class_name":"put api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Get Container Host","permalink":"/docs/1backend/get-host"},"next":{"title":"Check if Container Image is Pullable","permalink":"/docs/1backend/image-pullable"}}');var n=i(74848),o=i(28453),r=i(53746),s=i.n(r),l=i(56518),c=i.n(l),d=i(99972),p=i.n(d),h=i(25342),u=i.n(h),m=(i(44215),i(82223),i(24861));const b={id:"build-image",title:"Build an Image",description:"Builds a Docker image with the specified parameters.",sidebar_label:"Build an Image",hide_title:!0,hide_table_of_contents:!0,api:"eJy9Vt1u2zYUfhWC14rkJCuw6WpJWwxZiyaIm6vEQI+pY5kNRTIkZcczBOwh9oR7kuGQUmzF7rDd5MaQqfPznZ/vo7a8Qi+ctEEazUt+2UpVeQbsgxGP6JhsoEa2lmHJwhKZtyjkQmLFLDhoMKDz+YN+0Lf41EqHPlp9E0YHkBrdiV+JMsYo5xT5G7PoGum9NDrnGTcWHVDqq4qXPJpckTXPuMOnFn24NNWGl1tOIVEHegRrlRTRrfjuCfaWe7HEBujJOgoaJPoXt+dwA2FJf8fFvt+9ZDJBV0aAYpZOgoknERTr4/CM4zM0ViEvORUQNpYefXBS17zLeBX7tpAKj+f8MHr/47Q7u1HO0fFBcg0NHqb8Ag0OiciCmUV8TrMNJpU4SqNrqZ9LBQF9OEzUpelIhxUv70c97jHMXnzM/DuKwDtyOrJqLI6b3aZh8/3AwbUYM3lrtE/jPJtMDuu7/sSz/74fB7i6jP90LOyVXoGSFft9ev3l/yQYLyA6Z9xe3r0eHkdyeojkTkMblsbJP7B6MyTvjvckoNOg2BTdCh37GGO+DaQu4x5F62TY8PJ+yy8RHLqLllh2P+to5aD2tJDvB/lh05WgXWwwLA0pjG1px2xkJi9GMlXIXnh8LM3HHK1TZFhEfi6ND+W7n8/OTzllG8BMqbhUzz6k1637urHIHnqTB84WRimzxorNNwyYtyCQga5YMI+oGYhEA7ZwpolsvfOpHvbZ1FIz1JU1UgcSIUnxlwgVuoGAJb/oVyZOYsdhsPITbmKLaTq3O5X9OLD/lWpGnXsta2Mh6lOOVKMjXAszhAMRtwMbkNRSDwr9r17qulUQnNG5MM0e+JsrNm2tNY7mlaawDMGWRXE6B/GIuiKHgh+oygVrdbqhejtW4QqVsQ3qwKyCsDCuYQvj2Ma0jl1csb2N9X//+VcjhTO0BFKgP5mDxyoOhlQykFx6AQqpKUoK1D52bMD9281ntjrPJyPUviyK9Xqd17rNjauL3s8XUFt1cp5P8mVoFNUS0DX+ejFN2XdF+zXUNbpcmiKaFDRQGaJWn16mQnnGaW9TGyb5eT45cSI/+4XiWuNDA3oPaVJf0Gy4b0dt3Ltv3/qToN9T2r/CKpCa8MdebnviphvnJRgRINYwyzhRlAy2W5rbnVNdR8dPLToSjVnGV+AkzKlx97MuG1hDXH/EzfBNoMMJ0ZU6CqpNtHmlal02eFwIgTb8q+2+BN3cfeUZn/efNY2pyMXBmu4+WPOSx8+iuIxkEM+2XIGuWyqy5CkkERh6odlRmxBlwwMVNbzSmz2Ar6Uh1UG/VNVRl+02CUfXvdinVz/0eNGjZE1znHVd9w/vkJ7n",sidebar_class_name:"put api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},N=void 0,M={},g=[];function k(e){const t={code:"code",p:"p",...(0,o.R)(),...e.components};return(0,n.jsxs)(n.Fragment,{children:[(0,n.jsx)(m.default,{as:"h1",className:"openapi__heading",children:"Build an Image"}),"\n",(0,n.jsx)(s(),{method:"put",path:"/container-svc/image",context:"endpoint"}),"\n",(0,n.jsx)(t.p,{children:"Builds a Docker image with the specified parameters."}),"\n",(0,n.jsxs)(t.p,{children:["Requires the ",(0,n.jsx)(t.code,{children:"container-svc:image:build"})," permission."]}),"\n",(0,n.jsx)(m.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,n.jsx)(c(),{parameters:void 0}),"\n",(0,n.jsx)(p(),{title:"Body",body:{content:{"application/json":{schema:{properties:{contextPath:{description:"ContextPath is the local path to the build context",example:".",type:"string"},dockerfilePath:{description:"DockerfilePath is the local path to the Dockerfile",example:"Dockerfile",type:"string"},name:{description:"Name is the name of the image to build",example:"nginx:latest",type:"string"}},required:["contextPath","name"],type:"object"}}},description:"Build Image Request",required:!0}}),"\n",(0,n.jsx)(u(),{id:void 0,label:void 0,responses:{200:{description:"OK",content:{"application/json":{schema:{type:"object"}}}},400:{description:"Invalid JSON",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}}}})]})}function y(e={}){const{wrapper:t}={...(0,o.R)(),...e.components};return t?(0,n.jsx)(t,{...e,children:(0,n.jsx)(k,{...e})}):k(e)}}}]);