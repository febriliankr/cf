import classNames from "classnames";
import React, { ReactElement } from "react";

interface Props {
  title?: string;
  children: ReactElement | ReactElement[] | string;
  className?: string;
}

function PageContainer({
  title,
  children,
  className = "",
}: Props): ReactElement {
  return (
    <div className={classNames("max-w-6xl mx-auto", className)}>{children}</div>
  );
}

export default PageContainer;
