import React, { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import ResponsiveAppBar from "../components/ResponsiveAppBar";
import { Box, Typography, Divider } from "@mui/material";
import { Destination } from "../interfaces/Destination";
import { Location } from "../interfaces/Location";
import { destinations, locations } from "../mock/Data";

export function DestinationPage() {
  const { id } = useParams<{ id: string }>();
  const [destination, setDestination] = useState<Destination | null>(null);
  const [location, setLocation] = useState<Location | null>(null);
  const nav = useNavigate();

  useEffect(() => {
    const foundDestination = destinations.find(
      (dest) => dest.id.toString() === id
    );
    if (foundDestination) {
      setDestination(foundDestination);
      const foundLocation = locations.find(
        (loc) => loc.id === foundDestination.locationId
      );
      if (foundLocation) {
        setLocation(foundLocation);
      }
    } else {
      console.error("Destination not found");
    }
  }, [id, nav]);

  if (!destination || !location) {
    return (
      <Typography variant="h4" align="center">
        Loading...
      </Typography>
    );
  } else {
    return (
      <Box
        sx={{
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
          justifyContent: "space-between",
          backgroundColor: "#f9f9f9",
          minHeight: "100vh",
          width: "100vw",
        }}
      >
        <ResponsiveAppBar />
        <Box
          sx={{
            borderRadius: "1em",
            overflow: "hidden",
            boxShadow: "0 10px 30px rgba(0,0,0,0.1)",
            width: "100%",
            maxWidth: "1080px",
          }}
        >
          <Box sx={{ width: "100%", overflow: "hidden" }}>
            <img
              src={destination.image_url}
              alt="Destination"
              style={{
                width: "100%",
                maxHeight: "500px",
                objectFit: "cover",
              }}
            />
          </Box>
          <Box
            sx={{
              p: 3,
              overflowY: "auto",
              maxHeight: "calc(100vh - 300px - 64px)",
            }}
          >
            <Typography
              variant="h4"
              component="h2"
              sx={{ textAlign: "center", mb: 2 }}
            >
              {destination.name}
            </Typography>
            <Divider sx={{ mb: 2 }} />
            <Typography variant="h6" component="h3" sx={{ mb: 1 }}>
              Location: {location.name}
            </Typography>
            <Typography variant="subtitle1" component="h3" sx={{ mb: 1 }}>
              Country: {location.country}
            </Typography>
            <Typography variant="body1" sx={{ mb: 2 }}>
              {destination.description}
            </Typography>
          </Box>
        </Box>
        <Box
          sx={{
            width: "100vw",
            textAlign: "center",
            py: 2,
            bgcolor: "#3874cb",
            color: "white",
            height: "40px",
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
    );
  }
}
