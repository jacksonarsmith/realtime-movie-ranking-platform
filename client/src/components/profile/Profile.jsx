import { Avatar, Box, Card, CardHeader, Divider, Typography } from "@mui/material"
import { useEffect, useState } from "react"
import axios from "axios"
import Cookies from "js-cookie"

const Profile = () => {
    const [profile, setProfile] = useState({});

    useEffect(() => {
        const fetchProfile = async () => {
            try {
                const token = Cookies.get("token");
                const response = await axios.get(`${import.meta.env.VITE_DEV_API}profile`, {
                    headers: {
                        Authorization: `Bearer ${token}`
                    }
                });
                setProfile(response.data);
            } catch (error) {
                console.error("Error fetching profile:", error);
            }
        };
        fetchProfile();
    }, []);

    const getInitials = (name) => {
      if (!name) return "";
      const names = name.split(" ");
      const initials = names.map((n) => n[0]).join("");
      return initials.toUpperCase();
    };

    return (
      <Box
        sx={{
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
          height: "100vh",
        }}
      >
        <Typography variant="h1">
          Your Profile
        </Typography>
        <Divider sx={{ width: "100%" }} />
        <Card 
          sx={{ 
            width: "50%",
            p: 2,
            mt: 2,
          }}
        >
          <CardHeader
            avatar={
              <Avatar sx={{ bgcolor: "primary.main" }} aria-label="name">
                {getInitials(profile.name)}
              </Avatar>
            }
            title={profile.name}
            subheader={profile.email}
          />
        </Card>
      </Box>
    )
}

export default Profile