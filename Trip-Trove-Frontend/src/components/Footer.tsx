import * as React from "react";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";

const Footer: React.FC = () => {
  return (
    <Box
      sx={{
        gridArea: "footer",
        width: "100vw",
        py: 2,
        bgcolor: "#3874cb",
        marginTop: { md: "1.5em" },
        height: "40px",
      }}
    >
      <Typography
        variant="h6"
        component="div"
        align="center"
        color="white"
        marginTop="10px"
      >
        Trip Trove, toate drepturile rezervate @ MPP (kill me please)
      </Typography>
    </Box>
  );
};

export default Footer;
