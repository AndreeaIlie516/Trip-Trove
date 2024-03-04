import * as React from "react";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import ResponsiveAppBar from "./ResponsiveAppBar";
import DestinationGrid from "./DestinationGrid";
import background from "../assets/view_background.jpeg";

const GridTemplateAreas: React.FC = () => {
  return (
    <Box
      sx={{
        width: "100vw",
        minHeight: "100vh"
      }}
    >
      <Box
        sx={{
          display: "grid",
          gridTemplateColumns: "0.3fr 3fr 3fr 0.3fr",
          gap: 1,
          gridTemplateRows: "auto",
          gridTemplateAreas: `
            "header header header header"
            "hero hero hero hero"
            "aside main main sidebar"
            "footer footer footer footer"
        `,
        }}
      >
        <Box sx={{ gridArea: "header" }}>
          <ResponsiveAppBar />
        </Box>
        <Box
          sx={{
            gridArea: "hero",
            backgroundImage: `url(${background})`,
            backgroundSize: "cover",
            backgroundPosition: "center",
            color: "white",
            display: "flex",
            alignItems: "center",
            marginTop: { md: "-0.55em" },
            height: "475px",
            justifyContent: "center",
          }}
        >
          <Typography
            variant="h2"
            component="div"
            fontSize="70px"
            border="1px"
            sx={{
              textShadow: "2px 2px 4px rgba(0, 0, 0, 0.5)",
              fontWeight: "bold",
            }}
          >
            Welcome to Trip Trove!
          </Typography>
        </Box>
        <Box sx={{ gridArea: "aside" }}></Box>
        <Box
          sx={{
            gridArea: "main",
            display: "flex",
            flexDirection: "column",
            justifyContent: "center",
            paddingTop: "20px",
          }}
        >
          <DestinationGrid />
        </Box>
        <Box sx={{ gridArea: "sidebar" }}></Box>
        <Box
          sx={{
            gridArea: "footer",
            bgcolor: "#3874cb",
            marginTop: { md: "1.5em" },
            height: "70px",
          }}
        >
          <Typography
            variant="h6"
            component="div"
            align="center"
            color="white"
            marginTop="20px"
          >
            Trip Trove, toate drepturile rezervate @ MPP (kill me please)
          </Typography>
        </Box>
      </Box>
    </Box>
  );
};

export default GridTemplateAreas;
