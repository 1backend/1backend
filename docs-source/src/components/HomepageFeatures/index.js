import React from "react";
import clsx from "clsx";
import styles from "./styles.module.css";

const FeatureList = [
  {
    title: "Language Agnostic",
    Svg: require("@site/static/img/undraw_docusaurus_mountain.svg").default,
    description: (
      <>
        Build in your language of choice. With a microservices architecture and
        OpenAPI support, you can reuse existing services and develop new ones in
        any programming language.
      </>
    ),
  },
  {
    title: "AI and Microservice Deployment Made Easy",
    Svg: require("@site/static/img/undraw_docusaurus_tree.svg").default,
    description: (
      <>
        Built on Docker and other container engines, OpenOrch streamlines
        the deployment of AI models and microservices, keeping things efficient
        and scalable.
      </>
    ),
  },
  {
    title: "Flexible and Unopinionated",
    Svg: require("@site/static/img/undraw_docusaurus_react.svg").default,
    description: (
      <>
        Use what works for you. OpenOrch is designed for flexibility,
        giving you the freedom to adopt or rewrite components to suit your
        needs.
      </>
    ),
  },
];

function Feature({ Svg, title, description }) {
  return (
    <div className={clsx("col col--4")}>
      <div className="text--center">
        <Svg className={styles.featureSvg} role="img" />
      </div>
      <div className="text--center padding-horiz--md">
        <h3>{title}</h3>
        <p>{description}</p>
      </div>
    </div>
  );
}

export default function HomepageFeatures() {
  return (
    <section className={styles.features}>
      <div className="container">
        <div className="row">
          {FeatureList.map((props, idx) => (
            <Feature key={idx} {...props} />
          ))}
        </div>
      </div>
    </section>
  );
}
