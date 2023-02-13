import { Outlet } from "react-router-dom";
import DenseAppBar from "./internal/Appbar";

export default function Layout() {
  return (
    <>
      <DenseAppBar></DenseAppBar>
      <Outlet />
    </>
  );
}
