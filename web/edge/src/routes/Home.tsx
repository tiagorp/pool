import { Box, Button, TextareaAutosize, Typography } from "@mui/material";
import { useState } from "react";

export default function Home() {
  const [val, setVal] = useState("");

  async function handleClick() {
    let val = "";

    try {
      const response = await fetch("http://localhost:8080/api/v1/ping");
      val = await response.json();
    } catch (error) {
      console.log(error);
      val = "something went wrong; check console";
    }

    setVal(val);
  }

  return (
    <>
      <Typography variant="h3">Hello World!</Typography>

      <Box>
        <Typography variant="body1">Let us build</Typography>
        <Button variant="contained" onClick={handleClick}>
          Ping!
        </Button>
      </Box>

      <Box marginTop={5}>
        <TextareaAutosize value={val}></TextareaAutosize>
      </Box>
    </>
  );
}
