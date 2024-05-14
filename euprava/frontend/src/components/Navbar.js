import React, { useState, useEffect } from "react";
import { Link } from "react-router-dom";

const Navbar = () => {
  const [registerDropdown, setRegisterDropdown] = useState(false);
  const [userRole, setUserRole] = useState("");

  useEffect(() => {
    const handleOutsideClick = (e) => {
      if (registerDropdown && !e.target.closest(".register-dropdown")) {
        setRegisterDropdown(false);
      }
     
      setUserRole(localStorage.getItem("userrole"));
    };

    document.body.addEventListener("click", handleOutsideClick);

    return () => {
      document.body.removeEventListener("click", handleOutsideClick);
    };
  }, [registerDropdown]);

  const toggleRegisterDropdown = () => {
    setRegisterDropdown(!registerDropdown);
  };

  const handleLogout = () => {
    localStorage.setItem("token", "");
    localStorage.setItem("userrole", "");
  };

  return (
    <nav className="navbar navbar-expand-lg navbar-dark bg-dark">
      <div className="container">
        <Link className="navbar-brand" to="/">
          Home
        </Link>
        <button className="navbar-toggler" type="button">
          <span className="navbar-toggler-icon"></span>
        </button>
        <div className={`collapse navbar-collapse`} id="navbarSupportedContent">
          <ul className="navbar-nav me-auto mb-2 mb-lg-0">
            <li className="nav-item">
              <Link className="nav-link" to="/login">
                Login
              </Link>
            </li>
            {userRole === "ADMIN" ? (
              <li className="nav-item">
                <Link className="nav-link" to="/register/university">
                  university
                </Link>
              </li>
            ) : null}
            {userRole === "ADMIN" ? (
              <li className="nav-item">
                <Link className="nav-link" to="/register/dorm">
                  Dorm
                </Link>
              </li>
            ) : null}
            {userRole === "ADMIN" ? (
              <li className="nav-item">
                <Link className="nav-link" to="/register/health-care">
                  Health Care
                </Link>
              </li>
            ) : null}
            {userRole === "ADMIN" ? (
              <li className="nav-item">
                <Link className="nav-link" to="/register/food">
                  Food
                </Link>
              </li>
            ) : null}          
            <li className="nav-item col-md-2">
              <button className="btn btn-dark nav-link" onClick={handleLogout}>
                Logout
              </button>
            </li>
          </ul>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
