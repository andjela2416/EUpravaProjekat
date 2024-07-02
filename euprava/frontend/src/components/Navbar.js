import React, { useState, useEffect } from "react";
import { Link } from "react-router-dom";

const Navbar = () => {
  const [userRole, setUserRole] = useState("");

  useEffect(() => {
    const handleOutsideClick = (e) => {
      setUserRole(localStorage.getItem("user_type"));
    };

    document.body.addEventListener("click", handleOutsideClick);

    return () => {
      document.body.removeEventListener("click", handleOutsideClick);
    };
  }, []);

  const handleLogout = () => {
    localStorage.clear();
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
            <li className="nav-item">
              <Link className="nav-link" to="/register/user">
                Napravite nalog
              </Link>
            </li>
            {userRole === "STUDENT" || userRole === "PROFESOR" || userRole === "ADMIN" ? (
              <li className="nav-item">
                <Link className="nav-link" to="/university">
                  Fakultet
                </Link>
              </li>
            ) : null}
            {userRole === "STUDENT" || userRole === "DOKTOR" || userRole === "DEZURAN" ? (
              <li className="nav-item">
                <Link className="nav-link" to="/health-care">
                  Zdravstvo
                </Link>
              </li>
            ) : null}
            {userRole === "STUDENT" || userRole === "ADMIN" ? (
              <li className="nav-item">
                <Link className="nav-link" to="/dorms">
                  Domovi
                </Link>
              </li>
            ) : null}
              {userRole === "STUDENT" || userRole === "KUVAR" ? (
              <li className="nav-item">
                <Link className="nav-link" to="/food">
                  Menza
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